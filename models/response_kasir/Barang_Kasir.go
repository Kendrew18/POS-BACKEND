package response_kasir

type Read_Store_Barang_Kasir_Response struct {
	Kode_store    string                       `json:"kode_store"`
	Nama_store    string                       `json:"nama_store"`
	Detail_barang []Read_Barang_Kasir_Response `json:"detail_barang"`
}

type Read_Barang_Kasir_Response struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Nama_barang_kasir string  `json:"nama_barang_kasir"`
	Nama_satuan       string  `json:"nama_satuan"`
	Jumlah_pengali    float64 `json:"jumlah_pengali"`
}
