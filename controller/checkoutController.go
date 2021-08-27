package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetCheckoutsController(c echo.Context) error {
	checkout, e := database.GetCheckouts()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   checkout,
	})
}

func InsertCheckoutsController(c echo.Context) error {
	checkout := model.Checkouts{}
	c.Bind(&checkout)

	if err := config.DB.Create(&checkout).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   checkout,
	})
}

func AddCartToCheckoutByIdCartController(c echo.Context) error {
	checkout := model.Checkouts{}
	c.Bind(&checkout)

	res, err := database.AddCArttoCheckoutByIdCart(checkout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, errUpdt := database.UpdateCartStatus(checkout.Id_cart, "checkout")
	if errUpdt != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
