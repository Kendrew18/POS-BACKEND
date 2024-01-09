package response_kasir

type Read_Stock_Kasir_Response struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Nama_barang_kasir string  `json:"nama_barang_kasir"`
	Nama_satuan       string  `json:"nama_satuan"`
	Jumlah            float64 `json:"jumlah"`
	Harga             int64   `json:"harga"`
}
