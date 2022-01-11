package models

import (
	"fmt"
	"time"

	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/utils"
)

const CarBookingTableName = "car_bookings"

type CarBooking struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	StartDate time.Time  `json:"start_date" gorm:"not null;index:idx_car_bookings_date_car_id,unique"`
	EndDate   time.Time  `json:"end_date" gorm:"not null;index:idx_car_bookings_date_car_id,unique"`
	Car       *Car       `json:"car"`
	CarID     int        `json:"car_id"  gorm:"index:idx_car_bookings_date_car_id,unique"`
}

func (booking *CarBooking) Validate() (map[string]interface{}, bool) {
	if time.Now().Unix() > booking.StartDate.Unix() {
		return utils.Message(false, "Booking start date can't be in past"), false
	}

	if booking.EndDate.Unix() <= booking.StartDate.Unix() {
		return utils.Message(false, "Booking end date must be bigger than start date"), false
	}

	car := &Car{}
	if booking.CarID > 0 {
		err := GetDB().Table(CarTableName).Where("id = ?", booking.CarID).First(car).Error
		if err != nil {
			return utils.Message(false, "Car is not exists"), false
		}
	} else {
		return utils.Message(false, "Car is not exists"), false
	}
	booking.Car = car

	var availability int64
	err := GetDB().Table(CarBookingTableName).Where("car_id = ? AND start_date >= ? AND ((? BETWEEN start_date AND end_date) OR (? BETWEEN start_date AND end_date))",
		car.ID,
		time.Now(),
		booking.StartDate,
		booking.EndDate).Limit(1).Count(&availability).Error

	if err != nil {
		fmt.Println(err)
		return utils.Message(false, "Failed to check car availability"), false
	}

	if availability > 0 {
		return utils.Message(false, "Car is not available"), false
	}

	return utils.Message(true, "success"), true
}

func (booking *CarBooking) Create() map[string]interface{} {

	if resp, ok := booking.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(booking).Preload("Car.Color").Error
	if err != nil {
		if config.Get().DebugMode {
			return utils.Message(false, err.Error())
		}
		return utils.Message(false, "Validation Error")
	}

	resp := utils.Message(true, "success")
	resp["result"] = map[string]uint{"booking_id": booking.ID}
	return resp
}

func GetCarBooking(id int) *CarBooking {
	booking := &CarBooking{}
	err := GetDB().Table(CarBookingTableName).Preload("Car").Preload("Car.Color").Where("id = ?", id).First(booking).Error
	if err != nil {
		return nil
	}
	return booking
}

func GetCarBookings(carId int) []*CarBooking {
	bookings := make([]*CarBooking, 0)

	err := GetDB().Table(CarBookingTableName).Preload("Car").Preload("Car.Color").Where("car_id = ?", carId).Find(&bookings).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bookings
}

func DeleteCarBooking(carID int, bookingID int) bool {
	return GetDB().Where("id = ? AND car_id = ?", bookingID, carID).Delete(&CarBooking{}).Error == nil
}
