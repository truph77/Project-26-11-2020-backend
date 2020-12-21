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

// GetUser will get a user from db
func GetUser(c echo.Context) error {
	db := db.GetDB()
	user := models.User{}
	id := c.Param("id")

	db.Where("id = ?", id).Find(&user)

	return c.JSON(http.StatusOK, user)
}

//CreateUser ...
func CreateUser(c echo.Context) error {
	db := db.GetDB()

	user := models.User{}

	error := c.Bind(&user)
	if error != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	if user.Username == "" || user.Password == "" {
		return c.JSON(http.StatusNoContent, "cant create!")
	}

	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

//UpdateUser ...
func UpdateUser(c echo.Context) error {
	db := db.GetDB()

	id := c.Param("id")
	user := models.User{}

	error := db.Where("id = ?", id).First(&user).Error
	if error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	c.Bind(&user)
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

//DeleteUser ...
func DeleteUser(c echo.Context) error {
	db := db.GetDB()

	id := c.Param("id")
	db.Where("id = ?", id).Delete(models.User{})

	return c.JSON(http.StatusOK, "deleted with id: "+id)
}
