package main

import (
	"fmt"
	"go-rest-api/api/controller"
	"go-rest-api/server"

	"github.com/gorilla/mux"
)

// rest-server main.go is cmd utility to start rest api server.
// rest api we have developed here using gorilla mux liberary
// in this default /ping endpoint provided to validate server
// is up and running (health check endpoint)
func main() {
	fmt.Println("✨ Starting rest api server...⌛")

	rc := controller.NewRestController()

	r := mux.NewRouter()
	r.HandleFunc("/ping", rc.PingHandler)

	server.Start(r)
}
