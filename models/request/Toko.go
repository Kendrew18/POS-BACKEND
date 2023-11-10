package request

type Input_Toko_Request struct {
	Co           int    `json:"co"`
	Kode_toko    string `json:"kode_toko"`
	Nama_toko    string `json:"nama_toko"`
	Alamat       string `json:"alamat"`
	Nomor_telpon string `json:"nomor_telpon"`
	Kode_gudang  string `json:"kode_gudang"`
}

type Read_Toko_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
