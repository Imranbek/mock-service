package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"mock-service/internal/api/jokes"
	"mock-service/internal/config"
	"mock-service/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.Server.JokeURL)

	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Server.Host + ":" + cfg.Server.Port

	log.Print("starting server")
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("sutting server down")
}
