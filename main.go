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
	hostPort := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s", hostPort)
	log.Fatal(http.ListenAndServe(hostPort,
		handlers.LoggingHandler(os.Stderr, http.FileServer(http.Dir("."))),
	))
}