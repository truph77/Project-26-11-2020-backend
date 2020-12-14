package models

import (
	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
}
