package response_kasir

type Read_Request_Barang_Kasir_Response struct {
	Kode_request_barang_kasir string                                      `json:"kode_request_barang_kasir"`
	Tanggal_request           string                                      `json:"tanggal"`
	Nama_gudang               string                                      `json:"nama_gudang"`
	Nama_store                string                                      `json:"nama_store"`
	Status                    int                                         `json:"status"`
	Detail_barang             []Read_Barang_Request_Barang_Kasir_Response `json:"detail_barang"`
}

type Read_Barang_Request_Barang_Kasir_Response struct {
	Kode_barang_request string  `json:"kode_barang_request"`
	Nama_barang_kasir   string  `json:"nama_barang_kasir"`
	Jumlah              float64 `json:"jumlah"`
}

type Read_Barang_Request_Kasir_Response struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Jumlah            float64 `json:"jumlah"`
	Jumlah_pengali    float64 `json:"jumlah_pengali"`
}

type Status_Request_Barang_Kasir_Response struct {
	Status      int    `json:"kode_barang_kasir"`
	Nama_status string `json:"nama_status"`
}
