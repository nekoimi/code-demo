package main

import (
	"./handler"
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
