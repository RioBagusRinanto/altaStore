package model

import (
	"database/sql"
	"time"
)

type Categories struct {
	Id_category   int    `json:"id_category" form:"id_category" gorm:"primarykey"`
	Nama_category string `json:"nama_kategori" form:"nama_kategori"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}
