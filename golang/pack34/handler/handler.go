package handler

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UploadHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		// 返回html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(responseWriter, "server error.")
			return
		}
		_, _ = io.WriteString(responseWriter, string(data))
	} else if request.Method == "POST" {
		// 接受文件保存到本地
		file, head, err := request.FormFile("file")
		if err != nil {
			log.Printf("Get file error : %v \n", err.Error())
			return
		}
		defer file.Close()

		newFile, err := os.Create("./storage/upload/" + head.Filename)
		if err != nil {
			log.Printf("Create file error : %v \n", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			log.Printf("Save file error : %v \n", err.Error())
			return
		}

		_, _ = io.WriteString(responseWriter, "Upload success.")
	}
}
