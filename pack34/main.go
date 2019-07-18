package main

import (
	"fmt"
	"golearn/pack34/handler"
	"net/http"
)

func main()  {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe("0.0.0.0:8010", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
