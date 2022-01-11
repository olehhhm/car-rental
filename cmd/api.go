package main

import (
	"github.com/olehhhm/car-rental/config"
	"github.com/olehhhm/car-rental/models"
	"github.com/olehhhm/car-rental/routes"
)

func main() {
	config.Init()
	conf := config.Get()

	models.Init(
		conf.DatabaseConfig.Username,
		conf.DatabaseConfig.Password,
		conf.DatabaseConfig.Host,
		conf.DatabaseConfig.Port,
		conf.DatabaseConfig.Name)

	routes.Init()
}
