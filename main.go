package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defaultHttpPort = 8000

func main() {
	port := defaultHttpPort
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("invalid port number (%s): %v\n", os.Args[1], err)
		}
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port),
		handlers.LoggingHandler(os.Stderr, http.FileServer(http.Dir("."))),
	))
}

