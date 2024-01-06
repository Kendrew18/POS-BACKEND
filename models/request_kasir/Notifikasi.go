package request_kasir

type Read_Notifikasi_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}

type Update_Jumlah_Minimal_Request struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Jumlah_minimal    float64 `json:"jumlah_minimal"`
}
