package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func routeRegister(path string) *mux.Router {
	r := mux.NewRouter().PathPrefix(path).Subrouter()
	r.HandleFunc("/hello", helloHandler)
	return r
}

//helloHandler handles / requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
