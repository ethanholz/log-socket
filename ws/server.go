package ws

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/taigrr/log-socket/logger"
)

// var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func LogSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("upgrade:", err)
		return
	}
	defer c.Close()
	lc := logger.CreateClient()
	defer lc.Destroy()
	lc.SetLogLevel(logger.LTrace)
	logger.Info("Websocket client attached.")
	for {
		logEvent := lc.Get()
		logJSON, err := json.Marshal(logEvent)
		err = c.WriteMessage(websocket.TextMessage, logJSON)
		if err != nil {
			logger.Error("write:", err)
			break
		}
	}
}
