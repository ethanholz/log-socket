package lambo_log_socket

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Level int

const (
	LTrace Level = iota
	LDebug
	LInfo
	LWarn
	LError
	LPanic
	LFatal
)

var logger = logrus.New()
var clients []Client

func init() {
	stderrClient := CreateClient()
	go stderrClient.logStdErr()
	logger.SetLevel(logrus.DebugLevel)
}

func (c *Client) logStdErr() {
	select {
	case entry := <-c.writer:
		if c.LogLevel >= entry.level {
			fmt.Println(entry.Output)
		}
	}
}

func CreateClient() *Client {
	var client Client
	client.writer = make(LogWriter, 10)
	clients = append(clients, client)
	return &client
}

func (c *Client) Destroy() error {
	return nil
}

func (c *Client) GetLogLevel() Level {
	return c.LogLevel
}

func createLog(e Entry) {
	for _, c := range clients {
		go func(c Client, e Entry) {
			c.writer <- e
		}(c, e)
	}
}

// SetLogLevel set log level of logger
func (c *Client) SetLogLevel(level Level) {
	c.LogLevel = level
}

// Trace prints out logs on trace level
func Trace(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Debug(args...)
}

// Debug prints out logs on debug level
func Debug(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Debug(args...)
}

// Info prints out logs on info level
func Info(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Info(args...)
}

// Warn prints out logs on warn level
func Warn(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Warn(args...)
}

// Error prints out logs on error level
func Error(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Error(args...)
}

// Fatal prints out logs on fatal level
func Fatal(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Fatal(args...)
}

// Panic prints out logs on panic level
func Panic(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(2)
	entry.Panic(args...)
}

// fileInfo for getting which line in which file
func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
