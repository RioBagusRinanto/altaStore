package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
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
	jumlah, _ := strconv.Atoi(c.QueryParam("jumlah"))

	//get product
	products, _ := database.GetProductById(productId)
	b := model.Products{}
	mapstructure.Decode(products, &b)
	if products == nil {
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"status": "produk tidak ditemukan",
		})
	} else if b.Stock < jumlah {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "jumlah produk melebihi stock",
		})
	}

	//get cart dengan id customer & status
	carts, _ := database.GetCartByIdCust(customer, "cart")
	if carts == nil {
		//insert cart baru
		cart := model.Carts{}
		cart.Id_customer = customer
		cart.Total = b.Harga_satuan * jumlah
		cart.Status_pesanan = "cart"

		cartRes, errCart := database.CreateCart(cart)
		if errCart != nil {
			fmt.Println("gagal buat cart")
		}
		c := model.Carts{}
		mapstructure.Decode(cartRes, &c)

		detailCart := model.CartDetails{}
		detailCart.Id_cart = c.Id_cart
		detailCart.Id_product = b.Id_product
		detailCart.Jumlah = jumlah
		//create cart details
		_, errCartDetail := database.CreateCartDetail(detailCart)
		if errCartDetail != nil {
			fmt.Println(errCartDetail)
		}
		//update stock product
		_, errUpdateProd := database.UpdateStockProduct(b.Id_product, (b.Stock - jumlah))
		if errUpdateProd != nil {
			fmt.Println(errUpdateProd)
		}
		fmt.Println("cart sukses dibuat")
	} else {
		c := model.Carts{}
		mapstructure.Decode(carts, &c)

		cd := model.CartDetails{}
		cdetails, _ := database.GetCartDetailByIdCart(c.Id_cart)
		mapstructure.Decode(cdetails, &cd)
		//get product yg ada update bila ada yang sama
		if productId == cd.Id_product {
			_, errUpdateCD := database.UpdateCartDetailProduc(cd.Id_cart, productId, (cd.Jumlah + jumlah))
			if errUpdateCD != nil {
				fmt.Println(errUpdateCD)
			}
		}

		//update stock product
		_, errUpdateProd := database.UpdateStockProduct(b.Id_product, (b.Stock - jumlah))
		if errUpdateProd != nil {
			fmt.Println(errUpdateProd)
		}
		fmt.Println("cart sukses diupdate")
	}
	listCartCust, _ := database.GetCartByIdCust(customer, "cart")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create cart",
		"data":   listCartCust,
	})
}
