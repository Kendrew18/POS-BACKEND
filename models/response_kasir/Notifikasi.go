package response_kasir

type Read_Notifikasi_Kasir_Response struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Nama_barang_kasir string  `json:"nama_barang_kasir"`
	Jumlah            float64 `json:"jumlah"`
	Jumlah_minimal    float64 `json:"jumlah_minimal"`
	Status            int     `json:"status"`
}
