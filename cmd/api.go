package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/routes"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Creating config value
	conf := config.New()

	// Setup routes
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", routes.Home)

	// Setup port
	http.ListenAndServe(":"+conf.ServerPort, router)
}
