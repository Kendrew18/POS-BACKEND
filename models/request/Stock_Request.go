package request

type Input_Barang_Request struct {
	Co                int    `json:"co"`
	Kode_stock        string `json:"kode_stock"`
	Nama_Barang       string `json:"nama_barang"`
	Harga_jual        int64  `json:"harga_jual"`
	Satuan            string `json:"satuan"`
	Kode_jenis_barang string `json:"kode_jenis_barang"`
	Kode_gudang       string `json:"kode_gudang"`
}

type Read_Barang_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
