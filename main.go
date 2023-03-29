package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/taigrr/log-socket/browser"
	logger "github.com/taigrr/log-socket/log"
	"github.com/taigrr/log-socket/ws"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

func generateLogs() {
	for {
		logger.Info("This is an info log!")
		logger.Trace("This is a trace log!")
		logger.Debug("This is a debug log!")
		logger.Warn("This is a warn log!")
		logger.Error("This is an error log!")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	defer logger.Flush()
	flag.Parse()
	http.HandleFunc("/ws", ws.LogSocketHandler)
	http.HandleFunc("/", browser.LogSocketViewHandler)
	go generateLogs()
	dLogger := log.Default()
	lsLogger := logger.Default()
	lsLogger.SubLoggers = append(lsLogger.SubLoggers, logger.EnrichLogger(dLogger))
	lsLogger.Println("This should be printed out twice!")
	logger.Fatal(http.ListenAndServe(*addr, nil))
}
