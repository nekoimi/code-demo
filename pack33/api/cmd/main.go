package main

import (
	"../handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	router.GET("/users", handlers.UserHandler)

	return router
}

func main() {
	var router = RegisterHandlers()
	http.ListenAndServe("0.0.0.0:8000", router)
}
