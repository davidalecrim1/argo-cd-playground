package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Mate! 🤠"))
	})

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})

	log.Println("Listening in port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
