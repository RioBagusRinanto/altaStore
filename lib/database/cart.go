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

func GetCartByIdCust(idCust int, status string) (interface{}, error) {
	var carts []model.Carts

	if e := config.DB.Where("id_customer = ?", idCust, "AND status_pesanan = ?", status).First(&carts).Error; e != nil {
		return nil, e
	}
	return carts, nil
}

func CreateCart(cart model.Carts) (interface{}, error) {
	if err := config.DB.Create(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
