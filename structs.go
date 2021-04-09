package lambo_log_socket

import "time"

type LogWriter chan Entry

type Client struct {
	LogLevel Level `json:"level"`
	writer   LogWriter
}

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Output    string    `json:"output"`
	File      string    `json:"file"`
	Level     string    `json:"level"`
	level     Level
}
