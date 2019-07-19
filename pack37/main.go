package main

import (
	template2 "html/template"
	"log"
	"net/http"
	"time"
)

func main()  {
	app := http.NewServeMux()
	app.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		template, err := template2.ParseFiles("templates/layout.html", "templates/index.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		template.Execute(responseWriter, "Hello World.")
	})
	server := &http.Server{
		Addr: ":8001",
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler: app,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
