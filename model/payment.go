package model

import (
	"database/sql"
	"time"
)

type Payments struct {
	Id_metode_pembayaran int    `json:"id_metode_pembayaran" form:"id_metode_pembayaran" gorm:"primarykey"`
	Nama_metode          string `json:"nama_metode" form:"nama_metode"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            sql.NullTime `gorm:"index"`
}
