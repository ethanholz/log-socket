package log

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func (l *Logger) SetInfoDepth(depth int) {
	l.FileInfoDepth = depth
}

// Trace prints out logs on trace level
func (l Logger) Trace(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
}

// Formatted print for Trace
func (l Logger) Tracef(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
}

// Debug prints out logs on debug level
func (l Logger) Debug(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "DEBUG",
		level:     LDebug,
	}
	createLog(e)
}

// Formatted print for Debug
func (l Logger) Debugf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "DEBUG",
		level:     LDebug,
	}
	createLog(e)
}

// Info prints out logs on info level
func (l Logger) Info(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
}

// Formatted print for Info
func (l Logger) Infof(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
}

// Info prints out logs on info level
func (l Logger) Notice(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "NOTICE",
		level:     LNotice,
	}
	createLog(e)
}

// Formatted print for Info
func (l Logger) Noticef(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "NOTICE",
		level:     LNotice,
	}
	createLog(e)
}

// Warn prints out logs on warn level
func (l Logger) Warn(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "WARN",
		level:     LWarn,
	}
	createLog(e)
}

// Formatted print for Warn
func (l Logger) Warnf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "WARN",
		level:     LWarn,
	}
	createLog(e)
}

// Error prints out logs on error level
func (l Logger) Error(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "ERROR",
		level:     LError,
	}
	createLog(e)
}

// Formatted print for error
func (l Logger) Errorf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "ERROR",
		level:     LError,
	}
	createLog(e)
}

// Panic prints out logs on panic level
func (l Logger) Panic(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
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

// Formatted print for panic
func (l Logger) Panicf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
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
func (l Logger) Fatal(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "FATAL",
		level:     LFatal,
	}
	createLog(e)
	Flush()
	os.Exit(1)
}

// Formatted print for fatal
func (l Logger) Fatalf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "FATAL",
		level:     LFatal,
	}
	createLog(e)
	Flush()
	os.Exit(1)
}
