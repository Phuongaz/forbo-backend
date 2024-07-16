package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/phuongaz/forbo/common"
	"github.com/phuongaz/forbo/models"
	"github.com/phuongaz/forbo/routers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = common.InitSQLDB()
	if err != nil {
		log.Fatal("Error connecting to database ", err.Error())
	} else {
		log.Println("Database connected successfully")
	}

	err = models.InitModelsMigrate()
	if err != nil {
		log.Fatal("Error migrating models ", err.Error())
	} else {
		log.Println("Migrate models successfully")
	}

	r := routers.InitRouters()
	r.Run(":8000")
}
