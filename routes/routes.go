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
	e.GET("/shipments", controller.GetShipmentsController)

	//insert data
	e.POST("/shipments", controller.InsertShipmentsController)
	e.POST("/carts", controller.InsertCartController)
	e.POST("/cartdetails", controller.InsertCartDetailController)
	e.POST("/categories", controller.InsertCategoryController)
	e.POST("/checkouts", controller.InsertCheckoutsController)
	e.POST("/customers", controller.InsertCustomerController)
	e.POST("/payments", controller.InsertPaymentController)
	e.POST("/products", controller.InsertProductController)
	e.POST("/sellers", controller.InsertSellerController)
	e.POST("/shipments", controller.InsertShipmentsController)

	//get by key
	e.GET("/shipment", controller.GetShipmentByIdController)
	return e
}
