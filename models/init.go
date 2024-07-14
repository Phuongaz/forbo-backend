package models

import (
	"fmt"
	"log"

	"github.com/phuongaz/forbo/common"
)

func InitModelsMigrate() {
	err := common.SQLDBUser.AutoMigrate(&UserModel{})
	if err != nil {
		log.Fatal("Error migrating user model")
		fmt.Println(err)
	}
	common.SQLDBFeed.AutoMigrate(&Feed{})
	if err != nil {
		log.Fatal("Error migrating feed model")
		fmt.Println(err)
	}
}
