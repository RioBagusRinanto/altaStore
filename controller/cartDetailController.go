package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetCartDetailsController(c echo.Context) error {
	cartDetails, e := database.GetCartDetails()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   cartDetails,
	})
}

func InsertCartDetailController(c echo.Context) error {
	cartDetails := model.CartDetails{}
	c.Bind(&cartDetails)

	if err := config.DB.Create(&cartDetails).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "success",
		"Data":   cartDetails,
	})
}
