package model

import (
	"database/sql"
	"time"
)

type Carts struct {
	Id_cart        int       `json:"Id_cart" form:"id_cart" gorm:"primarykey"`
	Total          int       `json:"Total" form:"total"`
	Status_pesanan string    `json:"Status_pesanan" form:"status_pesanan"`
	Id_customer    int       `json:"Id_customer" form:"id_customer"`
	Customer       Customers `gorm:"foreignKey:Id_customer"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
