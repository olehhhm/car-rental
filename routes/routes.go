package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olehhhm/car-rental/config"
)

func Init() {
	// Setup routes
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", GetHome)
	router.Route("/car", CarRoute)

	// Setup port
	http.ListenAndServe(":"+config.Get().ServerPort, router)
}
