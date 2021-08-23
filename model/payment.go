package model

type Payments struct {
	Id   string `json:"id_metode_pembayaran" form:"id_metode_pembayaran"`
	Nama string `json:"nama_metode" form:"nama_metode"`
}
