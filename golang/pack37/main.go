package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main()  {
	app := http.NewServeMux()
	app.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		tpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = tpl.Execute(responseWriter, "Hello World."+request.Method)
		if err != nil {
			log.Println(err.Error())
			return
		}
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
