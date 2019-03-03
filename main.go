package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Flags is a structure for app arguments
type Flags struct {
	ListenPort int
}

func parseFlags() Flags {
	var f Flags
	flag.Parse()
	if flag.NArg() < 1 {
		f.ListenPort = 8080
	}
	port, err := strconv.Atoi(flag.Arg(0))
	if err != nil || port == 0 {
		return f
	}
	f.ListenPort = port
	return f
}

//HomeHandler handles / requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	flags := parseFlags()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	fmt.Println("Start serving on port", flags.ListenPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(flags.ListenPort), r))
}
