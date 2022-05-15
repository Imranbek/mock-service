package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"mock-service/internal/handler"
)

func main() {

	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("sutting server down")
}