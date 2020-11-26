package main

import (
	"github.com/labstack/echo"
	"github.com/truph77/handler"
)

func main() {
	e := echo.New()

	// Routes
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":7777"))
}
