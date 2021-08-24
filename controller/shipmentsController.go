package controller

import (
	"altaStore/lib/database"
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
