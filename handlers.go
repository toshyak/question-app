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

type addQuestionRequest struct {
	Question string `json:"question"`
}

type addQuestionResponse struct {
	QuestionID int    `json:"id"`
	Message    string `json:"message"`
}

type question struct {
	id           int
	questionText string
}

var questions []question

func routeRegister(router *mux.Router, path string) *mux.Router {
	r := router.PathPrefix(path).Subrouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/error", errorHandler).Methods("GET")
	r.HandleFunc("/addQuestion", addQuestion).Methods("POST")
	return r
}

//helloHandler handles / requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func addQuestion(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var questionRequest addQuestionRequest
	err := d.Decode(&questionRequest)
	if err != nil || questionRequest.Question == "" {
		fmt.Println(questionRequest)
		returnError(w, "Request could not be parsed", http.StatusBadRequest)
		return
	}
	var q question
	q.questionText = questionRequest.Question
	q.id = getMaxQuestionID() + 1
	questions = append(questions, q)
	fmt.Println(questions)
	w.Header().Set("Content-Type", "application/json")
	resp := addQuestionResponse{QuestionID: q.id, Message: "OK"}
	json.NewEncoder(w).Encode(resp)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	returnError(w, "Request failed", http.StatusInternalServerError)
}

func returnError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	e := errorResponse{Message: message}
	json.NewEncoder(w).Encode(e)
}

func getMaxQuestionID() int {
	max := 0
	for _, q := range questions {
		if max < q.id {
			max = q.id
		}
	}
	return max
}
