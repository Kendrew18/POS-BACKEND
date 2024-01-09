package response

type Kode_Kasir_Kode_Store_Response struct {
	Kode_kasir string `json:"kode_kasir"`
	Kode_store string `json:"kode_store"`
}

type Read_Request_Barang_Kasir_Stock_Response struct {
	Kode_request_barang_kasir string                                            `json:"kode_request_barang_kasir"`
	Tanggal_request           string                                            `json:"tanggal"`
	Kode_toko                 string                                            `json:"kode_toko"`
	Nama_toko                 string                                            `json:"nama_toko"`
	Jumlah                    float64                                           `json:"jumlah"`
	Status                    int                                               `json:"status"`
	Detail_barang             []Read_Barang_Request_Barang_Kasir_Stock_Response `json:"detail_barang"`
}

type Read_Barang_Request_Barang_Kasir_Stock_Response struct {
	Kode_barang_request string  `json:"kode_barang_request"`
	Kode_stock_gudang   string  `json:"kode_stock_gudang"`
	Nama_barang         string  `json:"nama_barang"`
	Jumlah              float64 `json:"jumlah"`
}
