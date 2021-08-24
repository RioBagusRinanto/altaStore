package model

type Checkouts struct {
	Id               int    `json:"Id_checkout" form:"id_checkout"`
	Total            int    `json:"Total" form:"total"`
	Status_transaksi string `json:"Status_transaksi" form:"status_transakasi"`
	Link_perjalanan  string `json:"Link_perjalanan" form:"link_perjalanan"`
	Id_shipment      int    `json:"Id_shipment" form:"id_shipment"`
	Id_cart          int    `json:"Id_cart" form:"id_cart"`
	Id_payment       int    `json:"Id_payment" form:"id_payment"`
}
