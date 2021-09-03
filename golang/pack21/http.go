package pack21

import (
	"net/http"
)

func Pack21HttpServer()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World."))
	})
	_ = http.ListenAndServe("0.0.0.0", nil)
}
