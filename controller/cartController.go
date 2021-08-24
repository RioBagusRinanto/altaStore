package controller

import (
	"altaStore/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetCartsController(c echo.Context) error {
	carts, e := database.GetCarts()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   carts,
	})
}
