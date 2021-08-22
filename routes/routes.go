package routes

import (
	"altaStore/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/products", controller.GetProductsController)

	return e
}
