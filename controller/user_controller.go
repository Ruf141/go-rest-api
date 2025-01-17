package controller

import (
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUseCase
}

func NewUserRepository(uu usecase.IUserUseCase) IUserController {
	return &userController{uu}
}
