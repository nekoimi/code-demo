package main

import (
	"fmt"
	. "golearn/pack36/conf"
	"golearn/pack36/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"html/template"
)

func demo () {
	res := NewMysql()
	fmt.Println(res.Db.Ping())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		content, err := ioutil.ReadFile("/home/developer/Golang/src/golearn/pack36/static/views/index.html")
		util.CheckError(err)
		_, err = responseWriter.Write(content)
		util.CheckError(err)
	})
	err := http.ListenAndServe("0.0.0.0:8001", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
