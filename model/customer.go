package model

type Customers struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Alamat   string `json:"alamat" form:"alamat"`
	Nama     string `json:"nama" form:"nama"`
}
