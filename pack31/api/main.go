package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()
	router.GET("/", func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, _ = io.WriteString(responseWriter, "hello.")
	})
	return router
}

func main()  {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
