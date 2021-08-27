package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	customer, _ := strconv.Atoi(c.QueryParam("iduser"))
	productId, _ := strconv.Atoi(c.QueryParam("idproduct"))
	// jumlah, _ := strconv.Atoi(c.QueryParam("jumlah"))
	// jumlah := json["jumlah"]

	//get product
	products, _ := database.GetProductById(productId)
	b, _ := json.Marshal(products)
	res := json.Unmarshal(b, &model.Products{})
	fmt.Println(res)
	// if products == nil {
	// 	return c.JSON(http.StatusNoContent, map[string]interface{}{
	// 		"status": "produk tidak ditemukan",
	// 	})
	// }

	//get cart dengan id customer & status
	carts, _ := database.GetCartByIdCust(customer, "cart")
	if carts == nil {
		//insert cart baru
		cart := model.Carts{}
		cart.Id_customer = customer
		cart.Total = 0
		cart.Status_pesanan = "cart"
		_, errCart := database.CreateCart(cart)
		if errCart != nil {
			fmt.Println("gagal buat cart")
		}
		fmt.Println("cart sukses dibuat")
	}
	listCartCust, _ := database.GetCarts()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create cart",
		"data":   listCartCust,
	})
	//get product id + validasi stok
	//
	return nil
}
