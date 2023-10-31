package request

type Input_Jenis_Barang_Request struct {
	Co                int    `json:"co"`
	Kode_jenis_barang string `json:"kode_jenis_barang"`
	Nama_Jenis_Barang string `json:"nama_jenis_barang"`
	Kode_gudang       string `json:"kode_gudang"`
}

type Read_Jenis_Barang_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
