package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Init(username string, password string, host string, port string, name string, debug bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)
	fmt.Println(dsn)

	var logg logger.Interface
	if debug {
		logg = logger.Default.LogMode(logger.Info)
	} else {
		logg = logger.Default.LogMode(logger.Silent)
	}
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logg,
	})
	if err != nil {
		panic(err)
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Car{}, &CarColor{}, &CarBooking{})
	db = conn
}
