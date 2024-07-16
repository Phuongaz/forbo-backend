package models

import (
	"github.com/phuongaz/forbo/common"
)

func InitModelsMigrate() error {
	err := common.SQLDBUser.AutoMigrate(&UserModel{})
	if err != nil {
		return err
	}
	common.SQLDBFeed.AutoMigrate(&Feed{})

	return nil
}
