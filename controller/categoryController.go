package controller

import (
	"altaStore/config"
	"altaStore/lib/database"
	"altaStore/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(c echo.Context) error {
	categories, e := database.GetCategories()
	// fmt.Println(products)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   categories,
	})
}

func InsertCategoryController(c echo.Context) error {
	categories := model.Categories{}
	c.Bind(&categories)

	if err := config.DB.Create(&categories).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   categories,
	})
}
