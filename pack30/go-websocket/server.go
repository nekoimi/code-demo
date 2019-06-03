package main

import (
	"./contract"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
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

	if _, ok := contract.Users[wsConn]; !ok {
		count := len(contract.Users) + 1
		contract.Users[wsConn] = "匿名用户_" + strconv.Itoa(count);
	}

	for  {
		data := conn.ReadMessage()
		conn.WriteMessage(data)
	}

}


func main()  {
	http.HandleFunc("/ws", webSocketHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
