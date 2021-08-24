package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetCheckouts() (interface{}, error) {
	var checkouts []model.Checkouts

	if e := config.DB.Find(&checkouts).Error; e != nil {
		return nil, e
	}
	return checkouts, nil

}
