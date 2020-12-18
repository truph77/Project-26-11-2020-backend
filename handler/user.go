package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/truph77/db"
	"github.com/truph77/models"
)

//GetUsers ...
func GetUsers(c echo.Context) error {
	db := db.GetDB()

	users := []models.User{}

	db.Find(&users) //select * from Login

	return c.JSON(http.StatusOK, users)
}

//CreateUser ...
func CreateUser(c echo.Context) error {
	db := db.GetDB()

	user := models.User{}
	if user.Password == "" || user.Username == "" {
		return c.String(http.StatusNoContent, "Cant create!")
	}
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

//UpdateUser ...
func UpdateUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

//DeleteUser ...
func DeleteUser(c echo.Context) error {
	db := db.GetDB()

	id := c.Param("id")
	// db.Where("id = ?", id).Delete(models.User{})
	db.Delete(models.User{}).Where("id = ?", id)

	return c.JSON(http.StatusOK, "deleted with id: "+id)
}
