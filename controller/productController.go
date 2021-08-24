package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, e := database.GetProducts()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   products,
	})
}

func InsertProductController(c echo.Context) error {
	products := model.Products{}
	c.Bind(&products)

	if err := config.DB.Create(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   products,
	})
}
