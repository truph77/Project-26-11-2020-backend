package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/truph77/db"
	"github.com/truph77/handler"
)

func main() {
	db.Init()
	defer db.CloseDB()

	e := echo.New()

	// Routes
	e.GET("/users", handler.GetUsers)
	e.GET("/user/:id", handler.GetUser)
	e.POST("/user", handler.CreateUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.POST("/api/login", handler.Login)

	e.Logger.Fatal(e.Start(":1234"))
}
