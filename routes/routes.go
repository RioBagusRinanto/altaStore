package routes

import (
	"altaStore/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/products", controller.GetProductsController)
	e.GET("/customers", controller.GetCustomersController)

	return e
}
