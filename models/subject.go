package models

import (
	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Country     string `json:"country"`
	Sponsor     string `json:"sponsor"`
}
