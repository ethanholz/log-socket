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

var logLevel Level
var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.DebugLevel)
}

// SetLogLevel set log level of logger
func SetLogLevel(level Level) {
	logLevel = level
}

// Trace prints out logs on trace level
func Trace(args ...interface{}) {
	if logLevel >= LTrace {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// Debug prints out logs on debug level
func Debug(args ...interface{}) {
	if logLevel >= LDebug {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// Info prints out logs on info level
func Info(args ...interface{}) {
	if logLevel >= LInfo {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// Warn prints out logs on warn level
func Warn(args ...interface{}) {
	if logLevel >= LWarn {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// Error prints out logs on error level
func Error(args ...interface{}) {
	if logLevel >= LError {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Fatal prints out logs on fatal level
func Fatal(args ...interface{}) {
	if logLevel >= LFatal {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// Panic prints out logs on panic level
func Panic(args ...interface{}) {
	if logLevel >= LPanic {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
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
