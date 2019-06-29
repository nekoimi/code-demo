package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func UserHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = io.WriteString(response, "hello world.")
}