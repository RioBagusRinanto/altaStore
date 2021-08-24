package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetPaymentsController(c echo.Context) error {
	payments, e := database.GetPayments()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   payments,
	})
}

func InsertPaymentController(c echo.Context) error {
	payments := model.Payments{}
	c.Bind(&payments)

	if err := config.DB.Create(&payments).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   payments,
	})
}
