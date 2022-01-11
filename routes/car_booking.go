package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/olehhhm/car-rental/models"
	"github.com/olehhhm/car-rental/utils"
)

func CarBookingRoute(route chi.Router) {
	route.Get("/", GetCarBookings)
	route.Post("/", CreateCarBooking)
	route.Route("/{bookingID}", func(r chi.Router) {
		r.Get("/", GetCarBooking)
		r.Delete("/", DeleteCarBooking)
	})
}

func GetCarBooking(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "bookingID")
	bookingID, err := strconv.Atoi(idParam)

	if err != nil {
		resp := utils.Message(false, idParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	booking := models.GetCarBooking(bookingID)
	if booking == nil {
		resp := utils.Message(false, "Record not exists")
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, resp)
		return
	}
	resp := utils.Message(true, "success")
	resp["result"] = booking
	utils.Respond(w, resp)
}

func GetCarBookings(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "success")

	idParam := chi.URLParam(r, "carID")
	carID, err := strconv.Atoi(idParam)

	if err != nil {
		resp := utils.Message(false, idParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	resp["result"] = models.GetCarBookings(carID)
	utils.Respond(w, resp)
}

func CreateCarBooking(w http.ResponseWriter, r *http.Request) {
	booking := &models.CarBooking{}
	err := json.NewDecoder(r.Body).Decode(booking)

	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	idParam := chi.URLParam(r, "carID")
	carID, err := strconv.Atoi(idParam)

	if err != nil {
		resp := utils.Message(false, idParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	booking.CarID = carID
	resp := booking.Create()
	utils.Respond(w, resp)
}

func DeleteCarBooking(w http.ResponseWriter, r *http.Request) {
	carIDParam := chi.URLParam(r, "carID")
	carID, err := strconv.Atoi(carIDParam)

	if err != nil {
		resp := utils.Message(false, carIDParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	bookingIDParam := chi.URLParam(r, "bookingID")
	bookingID, err := strconv.Atoi(bookingIDParam)

	if err != nil {
		resp := utils.Message(false, bookingIDParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	var resp map[string]interface{}
	if models.DeleteCarBooking(carID, bookingID) {
		resp = utils.Message(true, "success")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		resp = utils.Message(false, "failed")
	}

	utils.Respond(w, resp)
}
