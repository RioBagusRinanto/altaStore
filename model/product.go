package model

import (
	"database/sql"
	"time"
)

type Products struct {
	Id_product   int        `json:"id_product" form:"id_product" gorm:"primarykey"`
	Nama         string     `json:"nama" form:"nama"`
	Harga_satuan int        `json:"harga_satuan" form:"harga_satuan"`
	Stock        int        `json:"stock" form:"stock"`
	Deskripsi    string     `json:"deskripsi" form:"deskripsi"`
	Id_category  int        `json:"id_category" form:"id_category"`
	Id_seller    int        `json:"id_seller" form:"id_seller"`
	Categori     Categories `gorm:"foreignKey:Id_category;"`
	Seller       Sellers    `gorm:"foreignKey:Id_seller;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
}
