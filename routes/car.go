package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CarRoute(route chi.Router) {
	route.Route("/color", CarColorRoute)
	route.Get("/", GetCar) // GET
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Luke, I am Your car."))
}
