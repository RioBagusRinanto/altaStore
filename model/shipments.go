package model

type Shipments struct {
	Id   string `json:"Id" form:"id_shipment"`
	Nama string `json:"Nama_ekspedisi" form:"nama_ekspedisi"`
}
