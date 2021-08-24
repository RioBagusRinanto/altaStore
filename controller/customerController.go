package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomersController(c echo.Context) error {
	customers, e := database.GetCustomers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   customers,
	})
}

func InsertCustomerController(c echo.Context) error {
	customers := model.Customers{}
	c.Bind(&customers)

	if err := config.DB.Create(&customers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   customers,
	})
}
