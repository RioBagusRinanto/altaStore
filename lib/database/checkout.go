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

func AddCArttoCheckoutByIdCart(entry model.Checkouts) (interface{}, error) {
	if e := config.DB.Create(&entry).Error; e != nil {
		return nil, e
	}
	return entry, nil
}
