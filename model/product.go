package model

import (
	"database/sql"
	"time"
)

type Products struct {
	Id_product   int      `json:"Id_product" form:"id_product" gorm:"primarykey"`
	Nama         string   `json:"Nama" form:"nama"`
	Harga_satuan int      `json:"Harga_satuan" form:"harga_satuan"`
	Stock        int      `json:"Stock" form:"stock"`
	Deskripsi    string   `json:"Deskripsi" form:"deskripsi"`
	Id_category  int      `json:"Id_category" form:"id_category"`
	Seller       string   `json:"Seller" form:"username_seller"`
	Categori     Category `gorm:"foreignKey:Id_category"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
}
