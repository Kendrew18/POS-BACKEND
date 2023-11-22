package response

type Read_Toko_Response struct {
	Kode_toko    string `json:"kode_toko"`
	Nama_toko    string `json:"nama_toko"`
	Alamat       string `json:"alamat"`
	Nomor_telpon string `json:"nomor_telpon"`
}

type Read_Dropdown_Nama_Toko_Response struct {
	Kode_toko string `json:"kode_toko"`
	Nama_toko string `json:"nama_toko"`
}
