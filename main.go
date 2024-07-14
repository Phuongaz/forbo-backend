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

	common.InitSQLDB()
	log.Println("Database connected")

	models.InitModelsMigrate()
	log.Println("Models migrated")

	r := routers.InitRouters()

	r.Run(":8000")
}
