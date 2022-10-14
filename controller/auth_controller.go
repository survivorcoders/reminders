package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"survivorcoders.com/reminders/Claims"
	"survivorcoders.com/reminders/Services"
	"survivorcoders.com/reminders/entity"
	"time"
)

type AuthenticationController struct {
	UserManager *Services.UserManager
}

func (a *AuthenticationController) SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if err := a.UserManager.SignIn(username, password); err != nil {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &Claims.JwtCustomClaims{
		Name:  "Jon Snow",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (a *AuthenticationController) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func (a *AuthenticationController) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (a *AuthenticationController) SignUp(c echo.Context) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := entity.User{Name: name, Username: username, Password: password}
	// Throws unauthorized error
	if err := a.UserManager.SignUp(&user); err != nil {
		return echo.ErrForbidden
	}

	return c.JSON(http.StatusOK, user)
}
