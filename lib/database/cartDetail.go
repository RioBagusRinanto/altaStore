package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetCartDetails() (interface{}, error) {
	var cartDetails []model.CartDetails

	if e := config.DB.Find(&cartDetails).Error; e != nil {
		return nil, e
	}
	return cartDetails, nil

}
