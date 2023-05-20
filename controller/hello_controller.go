package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHelloController interface {
	SayHello(c echo.Context) error
}

type HelloController struct {
}

func NewHelloController() IHelloController {
	return &HelloController{}
}

func (hc *HelloController) SayHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
