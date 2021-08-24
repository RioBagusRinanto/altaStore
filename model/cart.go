package model

type Carts struct {
	Id_cart        int    `json:"Id_cart" form:"id_cart"`
	Total          int    `json:"Total" form:"total"`
	Status_pesanan string `json:"Status_pesanan" form:"status_pesanan"`
	Username       string `json:"Username" form:"username"`
}
