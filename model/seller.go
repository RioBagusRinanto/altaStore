package model

type Sellers struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Alamat   string `json:"alamat" form:"alamat"`
	Toko     string `json:"nama_toko" form:"nama_toko"`
	Email    string `json:"email" form:"email"`
}
