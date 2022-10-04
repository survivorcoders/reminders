package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"survivorcoders.com/reminders/controller"
	"survivorcoders.com/reminders/repository"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	dbConnection, err := gorm.Open(postgres.Open("postgres://localhost:5432/reminders"), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your CockroachDB cluster.")

	// Routes
	reminderController := controller.ReminderController{
		ReminderRepository: repository.ReminderRepository{
			DB: dbConnection,
		},
	}
	e.GET("/reminders", reminderController.GetAll)
	e.POST("/reminders", reminderController.Create)
	e.GET("/reminders/:id", reminderController.Get)
	e.PUT("/reminders/:id", reminderController.PUT)
	e.DELETE("/reminders/:id", reminderController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
