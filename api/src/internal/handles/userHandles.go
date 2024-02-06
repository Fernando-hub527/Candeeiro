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

func (userHandles *UserHandles) Login(ctx echo.Context) error {
	user, err := dtos.FactoryNewDTO(ctx.Request().Body, dtos.UserLoginDTO{})

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	registeredUser, err := userHandles.userCase.ValidLogin(user.UserName, user.Password, ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, errToken := jwt.CreateToken(registeredUser.UserName)
	if errToken != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, dtos.ResponseTokenDto{AccessToken: accessToken})
	return nil

}

func (userHandles *UserHandles) CreateUser(ctx echo.Context) error {
	user, err := dtos.FactoryNewDTO(ctx.Request().Body, dtos.NewUserDTO{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	newUser, err := userHandles.userCase.CreateUser(*user, ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, dtos.ResponseUserDTO{UserName: newUser.UserName, Email: newUser.Email, Telephone: newUser.Telephone})
	return nil

}
