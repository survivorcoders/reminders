package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"survivorcoders.com/reminders/Claims"
	"survivorcoders.com/reminders/Services"
	"survivorcoders.com/reminders/controller"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//database testing
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = database.AutoMigrate(&entity.Reminder{}, &entity.User{})
	if err != nil {
		return
	}

	reminderService := repository.RemindersProviderRepository{Database: database}
	authService := repository.UserProviderRepository{Database: database}
	userManager := Services.UserManager{UserRepository: &authService}
	reminderController := controller.ReminderController{ReminderService: &reminderService}
	authController := controller.AuthenticationController{UserManager: &userManager}

	//authentication routes
	e.POST("/sign-in", authController.SignIn)
	e.POST("/sign-up", authController.SignUp)

	// Unauthenticated route
	e.GET("/", authController.Accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims: &Claims.JwtCustomClaims{
			Admin: true,
		},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", authController.Restricted)

	// Routes
	r.GET("/reminders", reminderController.GetAllReminders)
	r.POST("/reminders", reminderController.PostCreateReminder)
	r.GET("/reminders/:id", reminderController.GetReminder)
	r.PUT("/reminders/:id", reminderController.PutUpdateReminder)
	r.DELETE("/reminders/:id", reminderController.DeleteReminder)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
