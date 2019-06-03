package contract

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Connection struct {
	wsConn *websocket.Conn
	in chan []byte
	out chan []byte
	Users map[*websocket.Conn]string  // 用户列表
}

func InitConnection (wsConn *websocket.Conn) *Connection  {

	conn := &Connection{
		wsConn: wsConn,
		in: make(chan []byte),
		out: make(chan []byte),
		Users: make(map[*websocket.Conn]string),
	}

	go conn.ReadLoop()
	go conn.WriteLoop()

	go func() {
		for  {
			conn.WriteMessage([]byte("heartbeat"))
			time.Sleep(30 * time.Second)
		}
	}()

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
			_ = conn.Close()
		}
		conn.in <- data
	}
}

func (conn *Connection) WriteLoop()  {
	for  {
		data := <- conn.out
		go func() {
			for keyConn := range Users {
				err := keyConn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Printf("WriteLoop Error : %v \n", err)
					_ = conn.Close()
				}
			}
		}()
	}
}
