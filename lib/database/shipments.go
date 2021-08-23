package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetShipments() (interface{}, error) {
	var shipments []model.Shipments

	if e := config.DB.Find(&shipments).Error; e != nil {
		return nil, e
	}
	return shipments, nil

}
