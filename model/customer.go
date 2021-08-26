package model

import (
	"database/sql"
	"time"
)

type Customers struct {
	Id_customer int    `json:"id_Customer" form:"id_customer" gorm:"primarykey"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Alamat      string `json:"alamat" form:"alamat"`
	Nama        string `json:"nama" form:"nama"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
}
