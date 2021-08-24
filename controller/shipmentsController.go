package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetShipmentsController(c echo.Context) error {
	shipments, e := database.GetShipments()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   shipments,
	})
}

func InsertShipmentsController(c echo.Context) error {
	shipments := model.Shipments{}
	c.Bind(&shipments)

	if err := config.DB.Create(&shipments).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   shipments,
	})
}
