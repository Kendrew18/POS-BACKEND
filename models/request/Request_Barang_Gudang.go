package request

type Update_Status_Request_Barang_Request struct {
	Status int `json:"status"`
}

type Update_Status_Kode_Request_Barang_Request struct {
	Kode_request_barang_kasir string `json:"kode_request_barang_kasir"`
	Kode_user                 string `json:"kode_user"`
	Kode_gudang               string `json:"kode_gudang"`
	Kode_nota                 string `json:"kode_nota"`
}
