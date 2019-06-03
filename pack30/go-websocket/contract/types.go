package contract

import "github.com/gorilla/websocket"

var (
	Users map[*websocket.Conn]string
)

func init()  {
	Users = make(map[*websocket.Conn]string)
}
