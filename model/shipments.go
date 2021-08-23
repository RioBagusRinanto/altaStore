package model

type Shipments struct {
	Id   string `json:"id" form:"id"`
	Nama string `json:"nama_ekspedisi" form:"nama_ekspedisi"`
}
