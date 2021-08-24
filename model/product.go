package model

// import "gorm.io/gorm"

type Products struct {
	// gorm.Model
	// Id        int    `json:"id" form:"id"`
	// Nama      string `json:"nama" form:"nama"`
	// Harga     int    `json:"harga_satuan" form:"harga_satuan"`
	// Stok      int    `json:"stock" form:"stock"`
	// Deskripsi string `json:"deskripsi" form:"deskripsi"`

	Id           int    `json:"Id_product" form:"id_product"`
	Nama         string `json:"Nama" form:"nama"`
	Harga_satuan int    `json:"Harga_satuan" form:"harga_satuan"`
	Stock        int    `json:"Stock" form:"stock"`
	Deskripsi    string `json:"Deskripsi" form:"deskripsi"`
	Id_category  int    `json:"Id_category" form:"id_category"`
	Seller       string `json:"Seller" form:"username_seller"`
	// IdKategori int
}
