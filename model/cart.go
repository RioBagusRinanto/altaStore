package model

import (
	"database/sql"
	"time"
)

type Carts struct {
	Id_cart        int       `json:"id_cart" form:"id_cart" gorm:"primarykey"`
	Total          int       `json:"total" form:"total"`
	Status_pesanan string    `json:"status_pesanan" form:"status_pesanan"`
	Id_customer    int       `json:"id_customer" form:"id_customer"`
	Customer       Customers `gorm:"foreignKey:Id_customer"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
