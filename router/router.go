package router

import (
	"asobi/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	hc controller.IHelloController,
	uc controller.IUserController,
) *echo.Echo {
	e := echo.New()

	e.GET("/hello", hc.SayHello)
	e.POST("/users", uc.StoreUser)
	return e
}
