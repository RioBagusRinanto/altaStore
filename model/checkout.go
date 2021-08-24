package model

import (
	"database/sql"
	"time"
)

type Checkouts struct {
	Id_checkout      int       `json:"Id_checkout" form:"id_checkout" gorm:"primarykey"`
	Total            int       `json:"Total" form:"total"`
	Status_transaksi string    `json:"Status_transaksi" form:"status_transakasi"`
	Link_perjalanan  string    `json:"Link_perjalanan" form:"link_perjalanan"`
	Id_shipment      int       `json:"Id_shipment" form:"id_shipment"`
	Id_cart          int       `json:"Id_cart" form:"id_cart"`
	Id_payment       int       `json:"Id_payment" form:"id_payment"`
	Shipment         Shipments `gorm:"foreignKey:Id_shipment"`
	Cart             Carts     `gorm:"foreignKey:Id_cart"`
	Payment          Payments  `gorm:"foreignKey:Id_payment"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
}
