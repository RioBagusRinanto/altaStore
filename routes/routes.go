package routes

import (
	"altaStore/constants"
	"altaStore/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//get all products
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

	e.POST("/sellers", controller.InsertSellerController)
	e.POST("/shipments", controller.InsertShipmentsController)

	//get by key
	e.GET("/shipment", controller.GetShipmentByIdController)
	e.GET("/product", controller.GetProductByIdController)

	//transaksi
	e.POST("/buy", controller.AddtoCartController)

	//seller
	e.GET("/sellersproduct", controller.GetProductBySellerController) //list produk dijual oleh seller
	e.POST("/products", controller.InsertProductController)           //entry produk baru
	e.PUT("/products", controller.UpdateProductByIsSeller)            //rubah detail produk
	e.DELETE("/product", controller.DeleteProductByIdController)      //hapus produk

	//customer
	e.GET("/products", controller.GetProductsController) //list produk
	e.GET("/carts/:userid", controller.GetCartByIdCustController)
	e.POST("/checkout", controller.AddCartToCheckoutByIdCartController)
	e.GET("/carts/:userid/:cartid", controller.GetDetailByIdCartController)
	e.POST("/login", controller.LoginCustController)

	//jwt auth
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/customer/:id", controller.GetUserDetailController)

	return e
}
