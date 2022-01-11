package models

import (
	"fmt"

	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/utils"
	"gorm.io/gorm"
)

const CarTableName = "cars"

type Car struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Color   *CarColor
	ColorID int
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
	resp["car"] = car
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
