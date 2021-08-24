package model

type Payments struct {
	Id   string `json:"Id_metode_pembayaran" form:"id_metode_pembayaran"`
	Nama string `json:"Nama_metode" form:"nama_metode"`
}
