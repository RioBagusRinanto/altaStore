package model

import (
	"database/sql"
	"time"
)

type Checkouts struct {
	Id_checkout      int       `json:"id_checkout" form:"id_checkout" gorm:"primarykey"`
	Total            int       `json:"total" form:"total"`
	Status_transaksi string    `json:"status_transaksi" form:"status_transakasi"`
	Link_perjalanan  string    `json:"link_perjalanan" form:"link_perjalanan"`
	Id_shipment      int       `json:"id_shipment" form:"id_shipment"`
	Id_cart          int       `json:"id_cart" form:"id_cart"`
	Id_payment       int       `json:"id_payment" form:"id_payment"`
	Shipment         Shipments `gorm:"foreignKey:Id_shipment"`
	Cart             Carts     `gorm:"foreignKey:Id_cart"`
	Payment          Payments  `gorm:"foreignKey:Id_payment"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
}
