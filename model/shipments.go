package model

import (
	"database/sql"
	"time"
)

type Shipments struct {
	Id_shipment    int    `json:"Id_shipment" form:"id_shipment" gorm:"primarykey"`
	Nama_ekspedisi string `json:"Nama_ekspedisi" form:"nama_ekspedisi"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
