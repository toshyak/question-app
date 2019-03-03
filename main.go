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

func routes() *mux.Router {
	router := routeRegister("/api/v1")
	return router
}

func main() {
	flags := parseFlags()
	r := routes()
	fmt.Println("Start serving on port", flags.ListenPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(flags.ListenPort), r))
}
