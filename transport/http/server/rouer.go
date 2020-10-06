package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Handlers []http.HandlerFunc

func Init() {
	Handlers = make([]http.HandlerFunc, 0)
	Handlers = append(Handlers, func(writer http.ResponseWriter, request *http.Request) {
		return
	})

	r := mux.NewRouter()
	for _, Handler := range Handlers {
		r.HandleFunc()
	}
}
