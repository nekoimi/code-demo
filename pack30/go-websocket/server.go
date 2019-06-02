package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Handler
func webSocketHandler(response http.ResponseWriter, request *http.Request)  {
	var (
		conn *websocket.Conn
		wsErr error
		data []byte
	)
	if conn, wsErr = upgrader.Upgrade(response, request, nil); wsErr != nil {
		return
	}

	for  {
		if _, data, wsErr = conn.ReadMessage(); wsErr != nil {
			goto WSERR
		}
		if wsErr = conn.WriteMessage(websocket.TextMessage, data); wsErr != nil {
			goto WSERR
		}
	}

WSERR:
	conn.Close()
}

func main()  {
	http.HandleFunc("/ws", webSocketHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
