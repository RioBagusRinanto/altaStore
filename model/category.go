package model

type Category struct {
	Id   string `json:"id" form:"id"`
	Name string `json:"nama_kategori" form:"nama_kategori"`
}
