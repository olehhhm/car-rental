package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/olehhhm/car-rental/models"
	"github.com/olehhhm/car-rental/utils"
)

func CarColorRoute(route chi.Router) {
	route.Get("/", GetCarColors)    // GET /articles
	route.Post("/", CreateCarColor) // POST /articles
}

func GetCarColors(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "success")
	resp["result"] = models.GetCarColors()
	utils.Respond(w, resp)
}

func CreateCarColor(w http.ResponseWriter, r *http.Request) {
	color := &models.CarColor{}
	err := json.NewDecoder(r.Body).Decode(color)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := color.Create()
	utils.Respond(w, resp)
}
