package log

import "time"

type (
	LogWriter chan Entry
	Level     int
)

const (
	LTrace Level = iota
	LDebug
	LInfo
	LNotice
	LWarn
	LError
	LPanic
	LFatal
)

type Client struct {
	LogLevel    Level `json:"level"`
	writer      LogWriter
	initialized bool
}

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Output    string    `json:"output"`
	File      string    `json:"file"`
	Level     string    `json:"level"`
	level     Level
}

func Default() *Logger {
	l := Logger{FileInfoDepth: 0, SubLoggers: []LevelLogger{}}
	return &l
}

func EnrichLogger(weak StdLogger) gapLogger {
	return gapLogger{subLogger: weak}
}

type Logger struct {
	FileInfoDepth int
	SubLoggers    []LevelLogger
}
type gapLogger struct {
	subLogger StdLogger
}
type (
	StdLogger interface {
		Println(v ...any)
		Printf(format string, v ...any)
		Print(v ...any)
		Panic(v ...any)
		Panicf(format string, v ...any)
		Panicln(v ...any)
		Fatal(v ...any)
		Fatalf(format string, v ...any)
		Fatalln(v ...any)
	}
	LevelLogger interface {
		Debug(v ...any)
		Debugf(format string, v ...any)
		Error(v ...any)
		Errorf(format string, v ...any)
		Errorln(v ...any)
		Fatal(v ...any)
		Fatalf(format string, v ...any)
		Fatalln(v ...any)
		Info(v ...any)
		Infof(format string, v ...any)
		Infoln(v ...any)
		Notice(v ...any)
		Noticef(format string, v ...any)
		Noticeln(v ...any)
		Panic(v ...any)
		Panicf(format string, v ...any)
		Panicln(v ...any)
		Print(v ...any)
		Printf(format string, v ...any)
		Println(v ...any)
		Trace(v ...any)
		Tracef(format string, v ...any)
		Traceln(v ...any)
		Warn(v ...any)
		Warnf(format string, v ...any)
		Warnln(v ...any)
	}
)
