package main

import (
	"fmt"
	"go-rest-api/server"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("✨ Starting rest api server...⌛")

	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)

	server.Start(r)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}
