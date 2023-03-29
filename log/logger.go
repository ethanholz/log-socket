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
	for _, sl := range l.SubLoggers {
		sl.Trace(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Tracef(format, args...)
	}
}

// Trace prints out logs on trace level with newline
func (l Logger) Traceln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "TRACE",
		level:     LTrace,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Traceln(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Debug(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Debugf(format, args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Info(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Infof(format, args...)
	}
}

// Info prints out logs on info level with newline
func (l Logger) Infoln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Infoln(args...)
	}
}

// Notice prints out logs on notice level
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
	for _, sl := range l.SubLoggers {
		sl.Notice(args...)
	}
}

// Formatted print for Notice
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
	for _, sl := range l.SubLoggers {
		sl.Noticef(format, args...)
	}
}

// Notice prints out logs on notice level with newline
func (l Logger) Noticeln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "NOTICE",
		level:     LNotice,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Noticeln(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Warn(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Warnf(format, args...)
	}
}

// Warn prints out logs on warn level with a newline
func (l Logger) Warnln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "WARN",
		level:     LWarn,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Warnln(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Error(args...)
	}
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

	for _, sl := range l.SubLoggers {
		sl.Errorf(format, args...)
	}
}

// Error prints out logs on error level with a new line
func (l Logger) Errorln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "ERROR",
		level:     LError,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Errorln(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Panic(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Panicf(format, args...)
	}
	panic(errors.New(output))
}

// Panic prints out logs on panic level with a newline
func (l Logger) Panicln(args ...interface{}) {
	output := fmt.Sprintln(args...)
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

	for _, sl := range l.SubLoggers {
		sl.Panicln(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Fatal(args...)
	}
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
	for _, sl := range l.SubLoggers {
		sl.Fatalf(format, args...)
	}
	os.Exit(1)
}

// Fatal prints fatal level with a new line
func (l Logger) Fatalln(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "FATAL",
		level:     LFatal,
	}
	createLog(e)
	Flush()
	for _, sl := range l.SubLoggers {
		sl.Fatalln(args...)
	}
	os.Exit(1)
}

// Handles print to info
func (l Logger) Print(args ...interface{}) {
	l.Info(args...)
	for _, sl := range l.SubLoggers {
		sl.Print(args...)
	}
}

// Handles formatted print to info
func (l Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
	for _, sl := range l.SubLoggers {
		sl.Printf(format, args...)
	}
}

// Handles print to info with new line
func (l Logger) Println(args ...interface{}) {
	output := fmt.Sprintln(args...)
	e := Entry{
		Timestamp: time.Now(),
		Output:    output,
		File:      fileInfo(l.FileInfoDepth),
		Level:     "INFO",
		level:     LInfo,
	}
	createLog(e)
	for _, sl := range l.SubLoggers {
		sl.Println(args...)
	}
}
