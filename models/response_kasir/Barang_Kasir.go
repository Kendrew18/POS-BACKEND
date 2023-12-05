package response_kasir

type Read_Barang_Kasir_Response struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Nama_barang_kasir string  `json:"nama_barang_kasir"`
	Nama_satuan       string  `json:"nama_satuan"`
	Jumlah_pengali    float64 `json:"jumlah_pengali"`
}
