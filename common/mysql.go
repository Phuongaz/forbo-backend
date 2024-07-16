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

func InitSQLDB() error {
	users_table := os.Getenv("DB_USERS_TABLE")
	feeds_table := os.Getenv("DB_FEEDS_TABLE")

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")

	log.Println("DB_HOST: ", db_host)
	log.Println("DB_PORT: ", db_port)
	log.Println("USERS TABLE: ", users_table)
	log.Println("FEEDS TABLE: ", feeds_table)

	dsnUser := "root:" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + users_table + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsnFeeds := "root:" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + feeds_table + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbUser, err := gorm.Open(mysql.Open(dsnUser), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	dbFeed, err := gorm.Open(mysql.Open(dsnFeeds), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	SQLDBUser = dbUser
	SQLDBFeed = dbFeed

	return nil
}
