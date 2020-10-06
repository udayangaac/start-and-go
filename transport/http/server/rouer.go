package server

import (
	transportHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var defaultServerOption = []transportHttp.ServerOption{
	transportHttp.ServerAfter(),
}

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", nil)
}
