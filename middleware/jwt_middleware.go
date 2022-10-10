package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func MubashirMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header["Authorization"]) > 0 {
			fmt.Println(c.Request().Header["Authorization"])

			//split and remove Bearer

			//read useremail or id from JWT claims

			//call user respository to get use using the email

			//c.Set("current_user_id",user.id)
		}
		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}
