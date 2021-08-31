package controller

import (
	"altaStore/lib/database"
	"altaStore/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	products, e := database.GetProducts()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   products,
	})
}

func InsertProductController(c echo.Context) error {
	products := model.Products{}
	c.Bind(&products)

	res, err := database.InsertProduct(products)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal simpan data",
			"error":  err,
		})
	} else if err == nil {
		res, _ = database.GetProductById(products.Id_product)

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func GetProductByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("idProduct"))
	// products := model.Products{}
	// c.Bind(&products)
	fmt.Println(id)
	var product, err = database.GetProductById(id)
	if product == nil && err == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   "produk tidak ditemukan / stok kosong",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   product,
	})
}

func GetProductBySellerController(c echo.Context) error {

	seller, _ := strconv.Atoi(c.QueryParam("id_seller"))
	product, _ := database.GetProductBySeller(seller)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   product,
	})

}

func UpdateProductByIsSeller(c echo.Context) error {
	product := model.Products{}
	c.Bind(&product)

	res, err := database.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal simpan data",
			"error":  err,
		})
	} else if err == nil {
		res, _ = database.GetProductById(product.Id_product)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func DeleteProductByIdController(c echo.Context) error {
	idProd, _ := strconv.Atoi(c.QueryParam("idproduct"))
	_, err := database.DeleteProductById(idProd)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
			"error":  err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "data sukses terhapus",
	})
}
