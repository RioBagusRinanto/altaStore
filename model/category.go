package model

type Category struct {
	Id_category   int    `json:"Id_category" form:"id_category"`
	Nama_category string `json:"Nama_kategori" form:"nama_kategori"`
}
