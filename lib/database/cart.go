package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetCarts() (interface{}, error) {
	var carts []model.Carts

	if e := config.DB.Find(&carts).Error; e != nil {
		return nil, e
	}
	return carts, nil

}
