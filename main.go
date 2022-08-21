package main

import (
	"log"
	"net/http"

	"example.com/app/storage"
)

const (
	ADDR = ":8080"
)

func main() {
	log.Printf("connect to DB")
	storage := storage.NewStorage()
	handler := NewHandlers(storage)

	mux := http.NewServeMux()

	mux.HandleFunc("/post", handler.Post)
	mux.HandleFunc("/get", handler.Get)

	log.Printf("server works on %s", ADDR)
	if err := http.ListenAndServe(ADDR, mux); err != nil {
		log.Fatal(err)
	}
}
