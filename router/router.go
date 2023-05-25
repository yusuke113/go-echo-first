package router

import (
	"asobi/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	uc controller.IUserController,
) *echo.Echo {
	e := echo.New()

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})
	e.POST("/users", uc.StoreUser)
	return e
}
