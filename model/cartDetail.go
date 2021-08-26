package model

import (
	"database/sql"
	"time"
)

type CartDetails struct {
	Id_cart    int      `json:"id_cart" form:"id_cart"`
	Id_product int      `json:"id_product" form:"id_product"`
	Jumlah     int      `json:"jumlah" form:"jumlah"`
	Cart       Carts    `gorm:"foreignKey:Id_cart"`
	Product    Products `gorm:"foreignKey:Id_product"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
}
