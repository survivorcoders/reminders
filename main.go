package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"survivorcoders.com/reminders/argonHash"
	"survivorcoders.com/reminders/controller"
	"survivorcoders.com/reminders/db"
	"survivorcoders.com/reminders/repository"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	db.Connect("postgres://postgres:admin@localhost:5432/reminders")

	// Establish the parameters to use for Argon2.
	var param = argonHash.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	// Routes
	reminderController := controller.ReminderController{
		ReminderRepository: repository.ReminderRepository{
			DB: db.Instance,
		},
	}
	authController := controller.AuthController{
		UserRepository: repository.UserRepository{
			DB: db.Instance,
		},
		Param: param,
	}

	//middleware (we validate JWT)

	//e.methode("path",<handler>,<middleware>)
	e.POST("/sign-up", authController.SignUp)
	e.POST("/sign-in", authController.SignIn)

	e.GET("/reminders", reminderController.GetAll)
	e.POST("/reminders", reminderController.Create)
	e.GET("/reminders/:id", reminderController.Get)
	e.PUT("/reminders/:id", reminderController.PUT)
	e.DELETE("/reminders/:id", reminderController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
