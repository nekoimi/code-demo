package main

import (
	"./contract"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Handler
func webSocketHandler(response http.ResponseWriter, request *http.Request) {
	var (
		wsConn *websocket.Conn
		wsErr  error
	)
	if wsConn, wsErr = upgrader.Upgrade(response, request, nil); wsErr != nil {
		return
	}

	conn := contract.InitConnection(wsConn)

	go func() {
		for  {
			conn.WriteMessage([]byte("heartbeat"))
			time.Sleep(time.Second)
		}
	}()

	for  {
		data := conn.ReadMessage()
		conn.WriteMessage(data)
	}

}


func main()  {
	http.HandleFunc("/ws", webSocketHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
