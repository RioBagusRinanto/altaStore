package model

type Shipments struct {
	Id_shipment    int    `json:"Id_shipment" form:"id_shipment"`
	Nama_ekspedisi string `json:"Nama_ekspedisi" form:"nama_ekspedisi"`
}
