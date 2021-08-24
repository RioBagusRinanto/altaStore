package model

import (
	"database/sql"
	"time"
)

type Sellers struct {
	Id_seller int    `json:"Id_seller" form:"id_seller" gorm:"primarykey"`
	Username  string `json:"Username" form:"username"`
	Password  string `json:"Password" form:"password"`
	Alamat    string `json:"Alamat" form:"alamat"`
	Nama_toko string `json:"Nama_toko" form:"nama_toko"`
	Email     string `json:"Email" form:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
