package main

import (
	"fmt"
	"go-rest-api/api/controller"
	"go-rest-api/server"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("✨ Starting rest api server...⌛")

	rc := controller.NewRestController()

	r := mux.NewRouter()
	r.HandleFunc("/ping", rc.PingHandler)

	server.Start(r)
}
