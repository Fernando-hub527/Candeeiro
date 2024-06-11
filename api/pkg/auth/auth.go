package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthenticateUser(next, errFUnc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName, err := ValidateToken(c.Request())
		if err != nil {
			return errFUnc(c)
		}
		c.SetParamNames(userName)
		return next(c)
	}
}

func SendAuthError(c echo.Context) error {
	return c.String(
		http.StatusAccepted,
		"access denied",
	)
}
