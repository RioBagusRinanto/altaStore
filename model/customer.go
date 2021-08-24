package model

type Customers struct {
	Username string `json:"Username" form:"username"`
	Email    string `json:"Email" form:"email"`
	Password string `json:"Password" form:"password"`
	Alamat   string `json:"Alamat" form:"alamat"`
	Nama     string `json:"Nama" form:"nama"`
}
