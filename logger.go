package lambo_log_socket

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

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
var clients []*Client
var sliceTex sync.Mutex

func init() {
	stderrClient := CreateClient()
	go stderrClient.logStdErr()
	logger.SetLevel(logrus.DebugLevel)
}

func (c *Client) logStdErr() {
	for {
		select {
		case entry, more := <-c.writer:
			if c.LogLevel >= entry.level {
				fmt.Println(entry.Output)
			}
			if !more {
				return
			}
		}
	}
}

func CreateClient() *Client {
	var client Client
	client.writer = make(LogWriter, 100)
	sliceTex.Lock()
	clients = append(clients, &client)
	sliceTex.Unlock()
	return &client
}

func (c *Client) Destroy() error {
	var otherClients []*Client
	sliceTex.Lock()
	c.writer = nil
	c = nil
	for _, x := range clients {
		if x != nil {
			otherClients = append(otherClients, x)
		}
	}
	clients = otherClients
	sliceTex.Unlock()
	return nil
}

func (c *Client) GetLogLevel() Level {
	return c.LogLevel
}

func createLog(e Entry) {
	sliceTex.Lock()
	for _, c := range clients {
		go func(c *Client, e Entry) {
			select {
			case c.writer <- e:
			default:
			}
		}(c, e)
	}
	sliceTex.Unlock()
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
