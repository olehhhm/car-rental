package models

import (
	"fmt"

	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/utils"
)

const CarColorTableName = "car_colors"

type CarColor struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique;not null"`
}

func (color *CarColor) Validate() (map[string]interface{}, bool) {

	if color.Name == "" {
		return utils.Message(false, "Color name should be exist"), false
	}

	return utils.Message(true, "success"), true
}

func (color *CarColor) Create() map[string]interface{} {

	if resp, ok := color.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(color).Error
	if err != nil {
		if config.Get().DebugMode {
			return utils.Message(false, err.Error())
		}
		return utils.Message(false, "Validation Error")
	}

	resp := utils.Message(true, "success")
	resp["result"] = color
	return resp
}

func GetCarColors() []*CarColor {
	colors := make([]*CarColor, 0)

	err := GetDB().Table(CarColorTableName).Find(&colors).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return colors
}
