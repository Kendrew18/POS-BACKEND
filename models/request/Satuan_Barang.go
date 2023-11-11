package request

type Input_Satuan_Barang_Request struct {
	Co                 int    `json:"co"`
	Kode_gudang        string `json:"kode_gudang"`
	Kode_satuan_barang string `json:"kode_satuan"`
	Nama_satuan_barang string `json:"nama_satuan"`
}

type Read_Satuan_Barang_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Delete_Satuan_Barang_Request struct {
	Kode_satuan_barang string `json:"kode_satuan_barang"`
}
