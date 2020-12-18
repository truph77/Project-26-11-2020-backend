package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/truph77/config"
	"github.com/truph77/models"
)

var db *gorm.DB
var err error

//Init used...
func Init() {
	config := config.GetConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_NAME, config.DB_PASSWORD)

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect successfully")
	}
	db.AutoMigrate(&models.User{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ...
func CloseDB() {
	db.Close()
}
