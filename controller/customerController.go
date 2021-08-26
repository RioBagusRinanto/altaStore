package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomersController(c echo.Context) error {
	customers, e := database.GetCustomers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   customers,
	})
}

func InsertCustomerController(c echo.Context) error {
	customers := model.Customers{}
	c.Bind(&customers)

	if err := config.DB.Create(&customers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   customers,
	})
}

func AddtoCartController(c echo.Context) error {
	var json map[string]interface{} = map[string]interface{}{}
	c.Bind(&json)
	customer := json["id_customer"].(int)
	productId := json["id_product"]
	jumlah := json["jumlah"]
	fmt.Println(customer)
	fmt.Println(productId)
	fmt.Println(jumlah)
	//get cart dengan id customer & status 1
	carts, _ := database.GetCartByIdCust(customer)
	if carts == nil {
		//insert cart baru
	}
	//get product id + validasi stok
	//
	return nil
}
