package handler

import (
	"github.com/labstack/echo"
	"github.com/truph77/models"
)

//GetUser used to get user info
func GetUser(c echo.Context) (err error) {
	//Connect DB
	// db, err := connectDB()
	// if err != nil {
	// 	return
	// }
	users := []models.User{}
}

//CreateUser ...
func CreateUser(c echo.Context) (err error) {

}

//UpdateUser...
func UpdateUser(c echo.Context) (err error) {

}

//DeleteUser ...
func DeleteUser(c echo.Context) (err error) {

}
