package controller

import (
	"asobi/model"
	"asobi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	StoreUser(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUseCase
}

func NewUserController(uu usecase.IUserUseCase) IUserController {
	return &userController{uu: uu}
}

func (uc *userController) StoreUser(c echo.Context) error {
	// POSTされたJSONをmodel.User型にデコードする
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザーを登録
	resUser, err := uc.uu.StoreUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resUser)
}
