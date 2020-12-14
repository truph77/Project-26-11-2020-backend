package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/truph77/models"
)

//GetUsers ...
func GetUsers(c echo.Context) error {
	// db := db.GetDB()
	users := []models.User{}
	fmt.Println(users)
	// db.Find(&users) //select * from Login

	return c.JSON(http.StatusOK, "users")
}

//CreateUser ...
func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

//UpdateUser ...
func UpdateUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

//DeleteUser ...
func DeleteUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}
