package common

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SQLDBUser *gorm.DB
var SQLDBFeed *gorm.DB

func InitSQLDB() {
	users_table := os.Getenv("USERS_TABLE")
	feeds_table := os.Getenv("FEEDS_TABLE")

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")

	dsnUser := "root:" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + users_table + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsnFeeds := "root:" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + feeds_table + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbUser, err := gorm.Open(mysql.Open(dsnUser), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to user database")
		fmt.Println(err)
	}

	log.Println("Connected to user database")

	dbFeed, err := gorm.Open(mysql.Open(dsnFeeds), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to user database")
		fmt.Println(err)
	}

	log.Println("Connected to feed database")

	SQLDBUser = dbUser
	SQLDBFeed = dbFeed
}
