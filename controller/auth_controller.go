package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"survivorcoders.com/reminders/dto"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
)

type AuthController struct {
	UserRepository repository.UserRepository
}

func (r AuthController) SignUp(c echo.Context) error {
	userEntity := &entity.User{}
	if err := c.Bind(userEntity); err != nil {
		return err
	}

	//generate the argon2 from the password sent
	//userEntity.Password = GenerateArgon2(userEntity.Password)
	if r.UserRepository.AddNewUser(userEntity) == nil {
		return c.JSON(http.StatusBadRequest, "Cannot create user..")
	}

	return c.JSON(http.StatusCreated, userEntity)
}

func (r AuthController) SignIn(c echo.Context) error {
	loginRequest := &dto.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ownerOfEmail := r.UserRepository.GetUserByEmail(loginRequest.Email)
	if ownerOfEmail == nil {
		return c.JSON(http.StatusBadRequest, "Cannot login ... ")
	}

	//// validate the password of this email with loginRequest.Password
	//if CompareArgon2(ownerOfEmail.Password, loginRequest.Password) != true {
	//	return c.JSON(http.StatusBadRequest, "Cannot login ... ")
	//}

	//generate a JWT token
	// using external libarary or using the echo library

	//send back the jwt as response hear
	c.Response().Header().Set("TOKEN", "GENERATED_JWT")

	return c.JSON(http.StatusCreated, ownerOfEmail)
}
