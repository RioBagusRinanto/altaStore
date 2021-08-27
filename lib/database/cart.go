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
	carts := model.Carts{}

	if e := config.DB.Preload("Customer").Where("id_customer = ? AND status_pesanan = ?", idCust, status).First(&carts).Error; e != nil {
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

func UpdatePriceTotal(idCart, total int) (interface{}, error) {
	cart := model.Carts{}
	if err := config.DB.Model(&cart).Where("id_cart = ?", idCart).Update("total", total).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func UpdateCartStatus(id int, status string) (interface{}, error) {
	cart := model.Carts{}
	if err := config.DB.Model(&cart).Where("id_cart = ?", id).Update("status_pesanan", status).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
