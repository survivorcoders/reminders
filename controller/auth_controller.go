package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"survivorcoders.com/reminders/argonHash"
	"survivorcoders.com/reminders/dto"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
	"time"
)

type AuthController struct {
	UserRepository repository.UserRepository
	Param          argonHash.Params
}

var (
	ErrCannotLogin  = errors.New("cannot login, problem with your email or password")
	ErrInCreateUser = errors.New("cannot create user,try again")
)

func (r AuthController) SignUp(c echo.Context) error {
	userEntity := &entity.User{}
	//Bind user Data to entity.User
	if err := c.Bind(userEntity); err != nil {
		return err
	}
	//generate the argon2 from the password sent
	if hashPass, err := argonHash.GenerateFromPassword(userEntity.Password, &r.Param); err != nil {
		return err
	} else {
		//userEntity.Password = GenerateArgon2(userEntity.Password)
		userEntity.Password = hashPass
	}
	//Add new user
	//true : return userEntityDATA
	//false : return string error
	if r.UserRepository.AddNewUser(userEntity) == nil {
		return c.JSON(http.StatusBadRequest, ErrInCreateUser)
	}

	return c.JSON(http.StatusCreated, userEntity)
}

func (r AuthController) SignIn(c echo.Context) error {
	//empty loginRequest data
	loginRequest := &dto.LoginRequest{}
	//bind data from client to loginRequest variable
	if err := c.Bind(loginRequest); err != nil {
		return err
	}
	//get a user (entity.User) data based on signIn email
	ownerOfEmail := r.UserRepository.GetUserByEmail(loginRequest.Email)
	if ownerOfEmail == nil {

		return c.JSON(http.StatusBadRequest, ErrCannotLogin)

	} else {

		// validate the password of this email sending by loginRequest.Password
		if res, _ := argonHash.ComparePasswordAndHash(loginRequest.Password, ownerOfEmail.Password); res == false {

			return c.JSON(http.StatusBadRequest, ErrCannotLogin)

		} else {
			//generate a JWT token
			token, err := GenerateJWT(ownerOfEmail.Name, ownerOfEmail.Email)
			if err != nil {
				return c.JSON(http.StatusBadRequest, ErrCannotLogin)
			}

			//send back the jwt as response header
			c.Response().Header().Set("TOKEN", token)

			// or cookie
			c.SetCookie(&http.Cookie{
				Name:    "token",
				Value:   token,
				Expires: time.Now().Add(time.Hour * 72).Local(),
			})
		}
	}

	return c.JSON(http.StatusCreated, ownerOfEmail)
}
