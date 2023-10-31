package response

type Read_Stock_Response struct {
	Kode_stock   string `json:"kode_stock"`
	Nama_Barang  string `json:"nama_barang"`
	Harga_jual   int64  `json:"harga_jual"`
	Jumlah       int    `json:"jumlah"`
	Satuan       string `json:"satuan"`
	Jenis_barang string `json:"jenis_barang"`
}

type Read_Barang_Response struct {
	Kode_jenis_barang  string                        `json:"kode_jenis_barang"`
	Jenis_Barang       string                        `json:"jenis_barang"`
	Read_Detail_Barang []Read_Detail_Barang_Response `json:"read_detail_barang"`
}

type Read_Detail_Barang_Response struct {
	Kode_stock  string `json:"kode_stock"`
	Nama_Barang string `json:"nama_barang"`
	Harga_jual  int64  `json:"harga_jual"`
	Satuan      string `json:"satuan"`
}
