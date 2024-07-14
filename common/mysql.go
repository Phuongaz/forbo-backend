package common

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SQLDBUser *gorm.DB
var SQLDBFeed *gorm.DB

func InitSQLDB() {

	dsnUser := "root:@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	dsnFeeds := "root:@tcp(127.0.0.1:3306)/feeds?charset=utf8mb4&parseTime=True&loc=Local"
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
