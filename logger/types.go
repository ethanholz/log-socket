package logger

import "time"

type LogWriter chan Entry
type Level int

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
