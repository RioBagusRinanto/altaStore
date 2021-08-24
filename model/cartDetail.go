package model

type CartDetails struct {
	Id_cart    int `json:"Id_cart" form:"id_cart"`
	Id_product int `json:"Id_product" form:"id_product"`
	Jumlah     int `json:"Jumlah" form:"jumlah"`
}
