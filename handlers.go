package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string `json:"message"`
}

func routeRegister(router *mux.Router, path string) *mux.Router {
	r := router.PathPrefix(path).Subrouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/error", errorHandler).Methods("GET")
	return r
}

//helloHandler handles / requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	e := errorResponse{Message: "Request Failed"}
	json.NewEncoder(w).Encode(e)
}
