package model

import (
	"database/sql"
	"time"
)

type Category struct {
	Id_category   int    `json:"Id_category" form:"id_category" gorm:"primarykey"`
	Nama_category string `json:"Nama_kategori" form:"nama_kategori"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}
