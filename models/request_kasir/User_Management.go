package request_kasir

type Input_User_Management struct {
	Co                   int    `json:"co"`
	Kode_user_management string `json:"kode_user_management"`
	Nama_store           string `json:"nama_store"`
}
