package lambo_log_socket

type LogWriter chan Entry

type Client struct {
	LogLevel Level `json:"level"`
	writer   LogWriter
}

type Entry struct {
	Timestamp string `json:"timestamp"`
	Output    string `json:"output"`
	Level     string `json:"level"`
	level     Level
}
