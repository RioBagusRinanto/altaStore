package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetCustomers() (interface{}, error) {
	var customers []model.Customers
	if e := config.DB.Find(&customers).Error; e != nil {
		return nil, e
	}

	return customers, nil

}
