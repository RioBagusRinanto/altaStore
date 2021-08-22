package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	// Id        int    `json:"id" form:"id"`
	// Nama      string `json:"nama" form:"nama"`
	// Harga     int    `json:"harga_satuan" form:"harga_satuan"`
	// Stok      int    `json:"stock" form:"stock"`
	// Deskripsi string `json:"deskripsi" form:"deskripsi"`

	Id           int    `json:"id_product" form:"id"`
	Nama         string `json:"nama" form:"nama"`
	Harga_satuan int    `json:"harga_satuan" form:"harga_satuan"`
	Stock        int    `json:"stock" form:"stock"`
	Deskripsi    string `json:"deskripsi" form:"deskripsi"`
	// IdKategori int
}
