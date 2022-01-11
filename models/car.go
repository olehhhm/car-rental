package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Color   CarColor
	ColorID int
}
