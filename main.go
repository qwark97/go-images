package main

import (
	"log"
	"net/http"

	"github.com/qwark97/go-images/handlers"
	"github.com/qwark97/go-images/storage"
)

const (
	ADDR = ":8080"
)

func main() {
	log.Printf("connect to DB")
	storage := storage.NewStorage()
	handler := handlers.NewHandlers(storage)

	mux := http.NewServeMux()

	mux.HandleFunc("/post", handler.Post)
	mux.HandleFunc("/get", handler.Get)

	log.Printf("server works on %s", ADDR)
	if err := http.ListenAndServe(ADDR, mux); err != nil {
		log.Fatal(err)
	}
}
