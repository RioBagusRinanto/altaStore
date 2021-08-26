package model

import (
	"database/sql"
	"time"
)

type Sellers struct {
	Id_seller int    `json:"id_seller" form:"id_seller" gorm:"primarykey"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Alamat    string `json:"alamat" form:"alamat"`
	Nama_toko string `json:"nama_toko" form:"nama_toko"`
	Email     string `json:"email" form:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
