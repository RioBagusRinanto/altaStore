package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetPayments() (interface{}, error) {
	var payments []model.Payments

	if e := config.DB.Find(&payments).Error; e != nil {
		return nil, e
	}
	return payments, nil

}
