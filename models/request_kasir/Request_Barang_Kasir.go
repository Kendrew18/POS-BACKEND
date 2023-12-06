package request_kasir

type Input_Request_Barang_Kasir_Request struct {
	Co                        int    `json:"co"`
	Kode_request_barang_kasir string `json:"kode_request_barang_kasir"`
	Kode_gudang               string `json:"kode_gudang"`
	Kode_store                string `json:"kode_store"`
	Kode_kasir                string `json:"kode_kasir"`
	Tanggal_request           string `json:"tanggal_request"`
	Status                    int    `json:"status"`
}

type Input_Barang_Request struct {
	Kode_stock_gudang string `json:"kode_stock_gudang"`
	Kode_barang_kasir string `json:"kode_barang_kasir"`
	Jumlah            string `json:"jumlah"`
}

type Input_Barang_Request_V2 struct {
	Co                        int     `json:"co"`
	Kode_request_barang_kasir string  `json:"kode_request_barang_kasir"`
	Kode_barang_request       string  `json:"kode_barang_request"`
	Kode_stock_gudang         string  `json:"kode_stock_gudang"`
	Kode_barang_kasir         string  `json:"kode_barang_kasir"`
	Jumlah                    float64 `json:"jumlah"`
}
