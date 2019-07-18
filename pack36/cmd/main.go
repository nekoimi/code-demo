package main

import (
	"fmt"
	. "golearn/pack36/conf"
	"time"
)

func main()  {
	res := NewMysql()
	fmt.Println(res.Db.Ping())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
