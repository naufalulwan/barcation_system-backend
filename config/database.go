package config

import (
	"barcation_be/helper"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=prefer TimeZone=Asia/Jakarta", ENV.DbHost, ENV.DbPort, ENV.DbName, ENV.DbUser, ENV.DbPass)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		helper.Logger.Errorf("\x1b[35mConnection Failed to Database\x1b[0m")
		log.Fatal("Connection Error : ", err)
	} else {
		helper.Logger.Infof("\x1b[38;5;10mConnection Success to Database\x1b[0m")
	}
}
