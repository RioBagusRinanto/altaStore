package config

import (
	"altaStore/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// config := map[string]string{
	// 	"DB_Username": "postgres",
	// 	"DB_Password": "postgres",
	// 	"DB_Port":     "5432",
	// 	"DB_Host":     "127.0.0.1",
	// 	"DB_Name":     "altastore",
	// }

	// connectioString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	"postgres",
	// 	"postgres",
	// 	"5432",
	// 	"localhost",
	// 	"altastore")
	connectioString := "host=localhost user=postgres password=postgres dbname=altastore port=5432"

	var e error

	DB, e = gorm.Open(postgres.Open(connectioString), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	InitMigrate()

}

func InitMigrate() {
	DB.AutoMigrate(&model.CartDetails{})
	DB.AutoMigrate(&model.Carts{})
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.Checkouts{})
	DB.AutoMigrate(&model.Customers{})
	DB.AutoMigrate(&model.Payments{})
	DB.AutoMigrate(&model.Products{})
	DB.AutoMigrate(&model.Sellers{})
	DB.AutoMigrate(&model.Shipments{})

}
