package handles

import (
	"net/http"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	jwt "github.com/Fernando-hub527/candieiro/internal/pkg/auth"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/labstack/echo/v4"
)

type UserHandles struct {
	userCase user.IUserUseCase
}

func NewUserHandles(userCase user.IUserUseCase) *UserHandles {
	return &UserHandles{
		userCase: userCase,
	}
}

func (*UserHandles) Login(ctx echo.Context) error {
	user, err := dtos.NewUserLogin(ctx.Param("password"), ctx.Param("userName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, errToken := jwt.CreateToken(user.UserName)
	if errToken != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, dtos.ResponseTokenDto{AccessToken: accessToken})
	return nil

}
