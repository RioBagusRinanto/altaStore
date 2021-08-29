package database

import (
	"altaStore/config"
	"altaStore/middlewares"
	"altaStore/model"
	"fmt"
)

func GetCustomers() (interface{}, error) {
	var customers []model.Customers
	if e := config.DB.Find(&customers).Error; e != nil {
		return nil, e
	}

	return customers, nil

}

func LoginCustomer(input model.Customers) (interface{}, error) {
	var e error
	if e = config.DB.Where("username = ? AND password =?", input.Username, input.Password).Find(&input).Error; e != nil {
		return nil, e
	}
	fmt.Println("find error - ", e)
	input.Token, e = middlewares.CreateToken(int(input.Id_customer))
	fmt.Println("token error - ", e)
	if e != nil {
		return nil, e
	}
	if e = config.DB.Save(input).Error; e != nil {
		return nil, e
	}

	return input, nil
}

func GetDetailCust(userId int) (interface{}, error) {
	var cust model.Customers
	fmt.Print(userId)
	// if e := config.DB.Where("id_customer = ?", userId).First(&cust).Error
	if e := config.DB.First(&cust, userId).Error; e != nil {
		return nil, e
	}
	fmt.Print(cust)

	return cust, nil
}
