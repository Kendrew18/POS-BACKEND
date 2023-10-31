package response

type Read_Stock_Response struct {
	Kode_stock   string `json:"kode_stock"`
	Nama_Barang  string `json:"nama_barang"`
	Harga_jual   int64  `json:"harga_jual"`
	Jumlah       int    `json:"jumlah"`
	Satuan       string `json:"satuan"`
	Jenis_barang string `json:"jenis_barang"`
}
