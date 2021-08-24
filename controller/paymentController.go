package controller

import (
	"altaStore/lib/database"
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
