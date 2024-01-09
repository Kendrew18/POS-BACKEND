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

type Input_Toko_Request_Barang_Request struct {
	Co           int    `json:"co"`
	Kode_toko    string `json:"kode_toko"`
	Nama_toko    string `json:"nama_toko"`
	Alamat       string `json:"alamat"`
	Nomor_telpon string `json:"nomor_telpon"`
	Kode_gudang  string `json:"kode_gudang"`
	Kode_kasir   string `json:"kode_kasir"`
	Kode_store   string `json:"kode_store"`
}

type Read_Request_Barang_Kasir_Stock_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Filter_Request_Barang_Kasir_Stock_Request struct {
	Tanggal_1 string `json:"tanggal_1"`
}
