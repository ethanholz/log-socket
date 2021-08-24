package log

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Trace prints out logs on trace level
func (_ Logger) Trace(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
}

// Formatted print for Trace
func (_ Logger) Tracef(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
}

// Debug prints out logs on debug level
func (_ Logger) Debug(args ...interface{}) {
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

// Formatted print for Debug
func (_ Logger) Debugf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
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
func (_ Logger) Info(args ...interface{}) {
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

// Formatted print for Info
func (_ Logger) Infof(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
}

// Info prints out logs on info level
func (_ Logger) Notice(args ...interface{}) {
	output := fmt.Sprint(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "NOTICE",
		level:     LNotice,
	}
	createLog(e)
}

// Formatted print for Info
func (_ Logger) Noticef(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(2),
		Level:     "NOTICE",
		level:     LNotice,
	}
	createLog(e)
}

// Warn prints out logs on warn level
func (_ Logger) Warn(args ...interface{}) {
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

// Formatted print for Warn
func (_ Logger) Warnf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
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
func (_ Logger) Error(args ...interface{}) {
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

// Formatted print for error
func (_ Logger) Errorf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
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
func (_ Logger) Panic(args ...interface{}) {
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

// Formatted print for panic
func (_ Logger) Panicf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
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
func (_ Logger) Fatal(args ...interface{}) {
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

// Formatted print for fatal
func (_ Logger) Fatalf(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
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
