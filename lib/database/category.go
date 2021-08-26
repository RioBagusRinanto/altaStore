package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetCategories() (interface{}, error) {
	var categories []model.Categories

	if e := config.DB.Find(&categories).Error; e != nil {
		return nil, e
	}
	return categories, nil

}
