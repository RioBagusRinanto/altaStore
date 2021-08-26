package database

import (
	"altaStore/config"
	"altaStore/model"
)

func GetProducts() (interface{}, error) {
	var products []model.Products

	if e := config.DB.Find(&products).Error; e != nil {
		return nil, e
	}
	// fmt.Println(products)
	return products, nil

}

func GetProductById(id int) (interface{}, error) {
	products := model.Products{}
	if err := config.DB.Preload("Categori").Preload("Seller").Find(&products, model.Products{Id_product: id}).Error; err != nil {
		return nil, err
	}
	if products.Stock == 0 {
		return nil, nil
	}
	return products, nil
}

func GetProductBySeller(id int) (interface{}, error) {
	products := []model.Products{}
	if err := config.DB.Preload("Seller").Preload("Categori").Find(&products, model.Products{Id_seller: id}).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func InsertProduct(prod model.Products) (interface{}, error) {

	if err := config.DB.Preload("Categori").Preload("Seller").Create(&prod).Error; err != nil {
		return nil, err
	}
	return prod, nil
}

func UpdateProduct(prod model.Products) (interface{}, error) {

	if err := config.DB.Set("gorm:association_autoupdate", false).Preload("Categori").Preload("Seller").Save(prod).Error; err != nil {
		return nil, err
	}
	return prod, nil

}
