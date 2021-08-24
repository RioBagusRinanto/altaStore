package model

type Category struct {
	Id   string `json:"Id" form:"id_category"`
	Name string `json:"Nama_kategori" form:"nama_kategori"`
}
