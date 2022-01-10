package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var conf *config.Config

func GetDB() *gorm.DB {
	return db
}

func GetConfig() *config.Config {
	return conf
}

func main() {
	initConfig()
	initDB()
	initRoutes()
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Creating config value
	conf = config.New()
}

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DatabaseConfig.Username,
		conf.DatabaseConfig.Password,
		conf.DatabaseConfig.Host,
		conf.DatabaseConfig.Port,
		conf.DatabaseConfig.Name,
	)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate()
	// db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func initRoutes() {
	// Setup routes
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", routes.Home)

	// Setup port
	http.ListenAndServe(":"+conf.ServerPort, router)
}
