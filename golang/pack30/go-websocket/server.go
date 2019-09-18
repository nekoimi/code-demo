package main

import (
	"./contract"
	"github.com/gorilla/websocket"
	"html/template"
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

func mainHandler(response http.ResponseWriter, request *http.Request)  {
	t, _ := template.ParseFiles("./message.html")
	_ = t.Execute(response, nil)
}


func main()  {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/ws", webSocketHandler)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}
