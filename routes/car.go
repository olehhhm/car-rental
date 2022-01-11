package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/olehhhm/car-rental/models"
	"github.com/olehhhm/car-rental/utils"
)

func CarRoute(route chi.Router) {
	route.Get("/", GetCars)
	route.Post("/", CreateCar)
	route.Get("/available", GetAvailableCarForBooking)
	route.Route("/{carID}", func(r chi.Router) {
		r.Get("/", GetCar)
		r.Delete("/", DeleteCar)
		r.Route("/booking", CarBookingRoute)
	})
	route.Route("/color", CarColorRoute)
}

func GetAvailableCarForBooking(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)

	dateLayout := time.RFC3339
	startDateParam := r.URL.Query().Get("start_date")
	startDate, err := time.Parse(dateLayout, startDateParam)

	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "start_date is not valid"))
		return
	}

	endDateParam := r.URL.Query().Get("end_date")
	endDate, err := time.Parse(dateLayout, endDateParam)
	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "end_date is not valid"))
		return
	}

	if time.Now().Unix() > startDate.Unix() {
		utils.Respond(w, utils.Message(false, "Booking start date can't be in past"))
		return
	}

	if endDate.Unix() <= startDate.Unix() {
		utils.Respond(w, utils.Message(false, "Booking end date must be bigger than start date"))
		return
	}

	resp := utils.Message(true, "success")
	resp["result"] = models.GetAvailableCars(startDate, endDate)
	utils.Respond(w, resp)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "carID")
	carID, err := strconv.Atoi(idParam)

	if err != nil {
		resp := utils.Message(false, idParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	car := models.GetCar(carID)
	if car == nil {
		resp := utils.Message(false, "Record not exists")
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, resp)
		return
	}
	resp := utils.Message(true, "success")
	resp["result"] = car
	utils.Respond(w, resp)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "carID")
	carID, err := strconv.Atoi(idParam)

	if err != nil {
		resp := utils.Message(false, idParam+" is not valid parameter")
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, resp)
		return
	}

	var resp map[string]interface{}
	if models.DeleteCar(carID) {
		resp = utils.Message(true, "success")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		resp = utils.Message(false, "failed")
	}

	utils.Respond(w, resp)
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "success")
	resp["result"] = models.GetCars()
	utils.Respond(w, resp)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	car := &models.Car{}
	err := json.NewDecoder(r.Body).Decode(car)

	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := car.Create()
	utils.Respond(w, resp)
}
