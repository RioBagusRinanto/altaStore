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

func CreateCartDetail(entry model.CartDetails) (interface{}, error) {

	if e := config.DB.Create(&entry).Error; e != nil {
		return nil, e
	}
	return entry, nil
}

func GetCartDetailByIdCart(id int) (interface{}, error) {
	carts := model.CartDetails{}

	if e := config.DB.Where("id_cart = ?", id).First(&carts).Error; e != nil {
		return nil, e
	}
	return carts, nil
}

func UpdateCartDetailProduc(idCart, IdProd, valProd int) (interface{}, error) {
	cd := model.CartDetails{}
	if e := config.DB.Model(cd).Where("id_cart = ? AND id_product = ?", idCart, IdProd).Update("jumlah", valProd).Error; e != nil {
		return nil, e
	}
	return cd, nil
}

func GetCartDetailIdcartIdprod(idcart, idprod int) (interface{}, error) {
	carts := model.CartDetails{}

	if e := config.DB.Where("id_cart = ? AND id_product = ?", idcart, idprod).First(&carts).Error; e != nil {
		return nil, e
	}
	return carts, nil
}

func GetCartDetailsByIdCart(id int) (interface{}, error) {
	carts := []model.CartDetails{}

	if e := config.DB.Where("id_cart = ?", id).Preload("Cart").Preload("Product").Find(&carts).Error; e != nil {
		return nil, e
	}
	return carts, nil
}
