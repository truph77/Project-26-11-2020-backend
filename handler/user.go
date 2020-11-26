package handler

import (
	"github.com/labstack/echo"
	"github.com/truph77/models"
)

//GetInfoUser used to get user info
func GetInfoUser(c echo.Context) (err error) {
	//Connect DB
	// db, err := connectDB()
	// if err != nil {
	// 	return
	// }
	users := []models.User{}

}
