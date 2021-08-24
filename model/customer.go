package model

import (
	"database/sql"
	"time"
)

type Customers struct {
	Id_customer int    `json:"Id_Customer" form:"id_customer" gorm:"primarykey"`
	Username    string `json:"Username" form:"username"`
	Email       string `json:"Email" form:"email"`
	Password    string `json:"Password" form:"password"`
	Alamat      string `json:"Alamat" form:"alamat"`
	Nama        string `json:"Nama" form:"nama"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
}
