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

func InsertShipments(s model.Shipments) (interface{}, error) {
	var shipments []model.Shipments

	if e := config.DB.Create(&s).Error; e != nil {
		return nil, e
	}

	return shipments, nil
}
