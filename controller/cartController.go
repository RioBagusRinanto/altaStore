package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetCartsController(c echo.Context) error {
	carts, e := database.GetCarts()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   carts,
	})
}

func InsertCartController(c echo.Context) error {
	carts := model.Carts{}
	c.Bind(&carts)

	if err := config.DB.Create(&carts).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   carts,
	})
}

func GetCartByIdCustController(c echo.Context) error {
	idcust, _ := strconv.Atoi(c.Param("userid"))

	carts, e := database.GetCartByIdCust(idcust, "cart")
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   carts,
	})
}
