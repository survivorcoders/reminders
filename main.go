package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"survivorcoders.com/reminders/controller"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	reminderController := controller.ReminderController{}
	e.GET("/reminders", reminderController.GetAll)
	e.POST("/reminders", reminderController.Create)
	e.GET("/reminders/:id", reminderController.Get)
	e.PUT("/reminders/:id", reminderController.PUT)
	e.DELETE("/reminders/:id", reminderController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
