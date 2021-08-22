package database

import (
	"altaStore/config"
	"altaStore/model"
	"fmt"
)

func GetProducts() (interface{}, error) {
	var products []model.Products

	if e := config.DB.Find(&products).Error; e != nil {
		return nil, e
	}
	fmt.Println(products)
	return products, nil

}
