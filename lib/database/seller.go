package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetSellers() (interface{}, error) {
	var sellers []model.Sellers

	if e := config.DB.Find(&sellers).Error; e != nil {
		return nil, e
	}
	return sellers, nil

}
