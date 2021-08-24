package routes

import (
	"altaStore/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//get all products
	e.GET("/products", controller.GetProductsController)
	e.GET("/customers", controller.GetCustomersController)
	e.GET("/sellers", controller.GetSellersController)
	e.GET("/payments", controller.GetPaymentsController)
	e.GET("/categories", controller.GetCategoriesController)
	e.GET("/carts", controller.GetCartsController)
	e.GET("/cartDetails", controller.GetCartDetailsController)
	e.GET("/checkouts", controller.GetCheckoutsController)

	return e
}
