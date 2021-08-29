package model

import (
	"database/sql"
	"time"
)

// status pesanan : cart -> dalam cart belum checkout, checkout -> sudah masuk checkout
type Carts struct {
	Id_cart        int           `json:"id_cart" form:"id_cart" gorm:"primarykey"`
	Total          int           `json:"total" form:"total"`
	Status_pesanan string        `json:"status_pesanan" form:"status_pesanan"`
	Id_customer    int           `json:"id_customer" form:"id_customer"`
	Customer       Customers     `gorm:"foreignKey:Id_customer"`
	Detail         []CartDetails `gorm:"foreignKey:Id_cart"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
