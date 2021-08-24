package model

type Carts struct {
	Id     string `json:"Id" form:"id_cart"`
	Total  int    `json:"Total" form:"total"`
	Status string `json:"Status_pemesanan" form:"status_pemesanan"`
	Buyer  string `json:"Username" form:"username"`
}
