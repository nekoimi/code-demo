package contract

import (
	"github.com/gorilla/websocket"
	"log"
)

type Connection struct {
	wsConn *websocket.Conn
	in chan []byte
	out chan []byte
}

func InitConnection (wsConn *websocket.Conn) *Connection  {
	conn := &Connection{
		wsConn: wsConn,
		in: make(chan []byte),
		out: make(chan []byte),
	}

	go conn.ReadLoop()
	go conn.WriteLoop()

	return conn
}

func (conn *Connection) ReadMessage () []byte  {
	data := <- conn.in
	return data
}

func (conn *Connection) WriteMessage (data []byte)  {
	conn.out <- data
}

func (conn *Connection) Close () error  {
	err := conn.wsConn.Close()
	return err
}

func (conn *Connection) ReadLoop()  {
	for  {
		_, data, err := conn.wsConn.ReadMessage()
		if err != nil {
			log.Printf("ReadLoop Error : %v \n", err)
			continue
		}
		conn.in <- data
	}
}

func (conn *Connection) WriteLoop()  {
	for  {
		data := <- conn.out
		err := conn.wsConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("WriteLoop Error : %v \n", err)
			continue
		}
	}
}
