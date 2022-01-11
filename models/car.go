package models

import (
	"fmt"
	"time"

	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/utils"
)

const CarTableName = "cars"

type Car struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name" gorm:"unique;not null"`
	Color     *CarColor  `json:"car_color"`
	ColorID   int        `json:"color_id"`
}

func (car *Car) Validate() (map[string]interface{}, bool) {

	if car.Name == "" {
		return utils.Message(false, "Car name should be exist"), false
	}
	carColor := &CarColor{}
	if car.ColorID > 0 {
		err := GetDB().Table(CarColorTableName).Where("id = ?", car.ColorID).First(carColor).Error
		if err != nil {
			return utils.Message(false, "Car color not exists"), false
		}
	} else {
		err := GetDB().Table(CarColorTableName).Where("name = ?", car.Color.Name).First(carColor).Error
		if err != nil {
			return utils.Message(false, "Car color not exists"), false
		}
	}
	car.Color = carColor
	return utils.Message(true, "success"), true
}

func (car *Car) Create() map[string]interface{} {

	if resp, ok := car.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(car).Error
	if err != nil {
		if config.Get().DebugMode {
			return utils.Message(false, err.Error())
		}
		return utils.Message(false, "Validation Error")
	}

	resp := utils.Message(true, "success")
	resp["result"] = car
	return resp
}

func GetCar(id int) *Car {
	car := &Car{}
	err := GetDB().Table(CarTableName).Preload("Color").Where("id = ?", id).First(car).Error
	if err != nil {
		return nil
	}
	return car
}

func GetAvailableCars(startDate time.Time, endDate time.Time) []*Car {
	bookings := make([]*CarBooking, 0)
	err := GetDB().
		Select("car_id").
		Table(CarBookingTableName).
		Where("start_date >= ? AND ((? BETWEEN start_date AND end_date) OR (? BETWEEN start_date AND end_date))",
			time.Now(),
			startDate,
			endDate).
		Group("car_id").
		Find(&bookings).
		Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(bookings) == 0 {
		return GetCars()
	}

	bookedCarIds := make([]int, 0)
	for _, book := range bookings {
		bookedCarIds = append(bookedCarIds, book.CarID)
	}

	cars := make([]*Car, 0)
	err = GetDB().
		Table(CarTableName).
		Preload("Color").
		Where("id NOT IN (?)", bookedCarIds).
		Find(&cars).
		Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cars
}

func GetCars() []*Car {
	cars := make([]*Car, 0)

	err := GetDB().Table(CarTableName).Preload("Color").Find(&cars).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cars
}

func DeleteCar(id int) bool {
	return GetDB().Delete(&Car{}, id).Error == nil
}
