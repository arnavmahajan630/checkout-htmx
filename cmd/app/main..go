package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	handler "github.com/arnavmahajan630/checkout-htmx/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	r := chi.NewMux()
	r.Handle("/public/*", public())
	r.Get("/", handler.Make(handler.Home))
	slog.Info("HTTP Server started", "Port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
