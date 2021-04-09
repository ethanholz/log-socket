package lambo_log_socket

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
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

var (
	clients        []*Client
	sliceTex       sync.Mutex
	stderrClient   *Client
	cleanup        sync.Once
	stderrFinished chan bool
)

func init() {
	stderrClient = CreateClient()
	stderrClient.SetLogLevel(LTrace)
	stderrFinished = make(chan bool, 1)
	go stderrClient.logStdErr()
}

func (c *Client) logStdErr() {
	for {
		select {
		case e, more := <-c.writer:
			if e.level >= c.LogLevel {
				fmt.Fprintf(os.Stderr, "%s\t%s\t%s\t%s\n", e.Timestamp.String(), e.Level, e.Output, e.File)
			}
			if !more {
				stderrFinished <- true
				return
			}
		}
	}
}

func CreateClient() *Client {
	var client Client
	client.initialized = true
	client.writer = make(LogWriter, 1000)
	sliceTex.Lock()
	clients = append(clients, &client)
	sliceTex.Unlock()
	return &client
}

func Flush() {
	cleanup.Do(func() {
		close(stderrClient.writer)
		<-stderrFinished
		stderrClient.Destroy()
	})
}

func (c *Client) Destroy() error {
	var otherClients []*Client
	if !c.initialized {
		panic(errors.New("Cannot delete uninitialized client, did you use CreateClient?"))
	}
	sliceTex.Lock()
	c.writer = nil
	c.initialized = false
	for _, x := range clients {
		if x.initialized {
			otherClients = append(otherClients, x)
		}
	}
	clients = otherClients
	sliceTex.Unlock()
	return nil
}

func (c *Client) GetLogLevel() Level {
	if !c.initialized {
		panic(errors.New("Cannot get level for uninitialized client, use CreateClient instead"))
	}
	return c.LogLevel
}

func createLog(e Entry) {
	sliceTex.Lock()
	for _, c := range clients {
		func(c *Client, e Entry) {
			select {
			case c.writer <- e:
				// try to clear out one of the older entries
			default:
				select {
				case <-c.writer:
					c.writer <- e
				default:
				}
			}
		}(c, e)
	}
	sliceTex.Unlock()
}

// SetLogLevel set log level of logger
func (c *Client) SetLogLevel(level Level) {
	if !c.initialized {
		panic(errors.New("Cannot set level for uninitialized client, use CreateClient instead"))
	}
	c.LogLevel = level
}

func (c *Client) Get() Entry {
	if !c.initialized {
		panic(errors.New("Cannot get logs for uninitialized client, did you use CreateClient?"))
	}
	return <-c.writer
}

// Trace prints out logs on trace level
func Trace(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
	//	entry := logger.WithFields(logrus.Fields{})
	//	entry.Data["file"] = fileInfo(2)
	//	entry.Debug(args...)
}

// Debug prints out logs on debug level
func Debug(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "DEBUG",
		level:     LDebug,
	}
	createLog(e)

}

// Info prints out logs on info level
func Info(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
}

// Warn prints out logs on warn level
func Warn(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "WARN",
		level:     LWarn,
	}
	createLog(e)
}

// Error prints out logs on error level
func Error(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "ERROR",
		level:     LError,
	}
	createLog(e)
}

// Panic prints out logs on panic level
func Panic(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "PANIC",
		level:     LPanic,
	}
	createLog(e)
	if len(args) >= 0 {
		switch args[0].(type) {
		case error:
			panic(args[0])
		default:
			// falls through to default below
		}
	}
	Flush()
	panic(errors.New(output))
}

// Fatal prints out logs on fatal level
func Fatal(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "FATAL",
		level:     LFatal,
	}
	createLog(e)
	Flush()
	os.Exit(1)
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
