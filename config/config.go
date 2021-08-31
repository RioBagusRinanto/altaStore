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
	connectioString := "host=ec2-3-230-61-252.compute-1.amazonaws.com user=vuetlwgfsgchav password=61d770f099a4553df03c805ff7e21301a11f608a98e2e8b8fd85c4ce257eb331 dbname=d1as2g2c8u0g3h port=5432"

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
	DB.AutoMigrate(&model.Categories{})
	DB.AutoMigrate(&model.Checkouts{})
	DB.AutoMigrate(&model.Customers{})
	DB.AutoMigrate(&model.Payments{})
	DB.AutoMigrate(&model.Products{})
	DB.AutoMigrate(&model.Sellers{})
	DB.AutoMigrate(&model.Shipments{})

}
