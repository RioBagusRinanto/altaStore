package model

type Sellers struct {
	Username string `json:"Username" form:"username"`
	Password string `json:"Password" form:"password"`
	Alamat   string `json:"Alamat" form:"alamat"`
	Toko     string `json:"Nama_toko" form:"nama_toko"`
	Email    string `json:"Email" form:"email"`
}
