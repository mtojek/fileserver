package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
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

	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("can't read network interfaces: %v", err)
	}

	log.Println("Listening on:")
	for _, iface := range interfaces {
		host := iface.String()
		if i := strings.Index(host, "/"); i > 0 {
			host = host[:i]
		}

		if !isIPv4(host) {
			continue
		}

		log.Printf("http://%s:%d/", host, port)
	}

	hostPort := fmt.Sprintf("0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(hostPort,
		handlers.LoggingHandler(os.Stderr, http.FileServer(http.Dir("."))),
	))
}

func isIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}
