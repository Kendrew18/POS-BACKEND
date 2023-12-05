package request_kasir

type Input_User_Management_Request struct {
	Co                   int    `json:"co"`
	Kode_user_management string `json:"kode_user_management"`
	Nama_store           string `json:"nama_store"`
	Kode_kasir           string `json:"kode_kasir"`
}

type Read_User_Management_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}
