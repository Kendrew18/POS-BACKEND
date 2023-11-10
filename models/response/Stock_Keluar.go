package response

type Read_Stock_Keluar_Response struct {
	Kode_stock_keluar                string                              `json:"kode_stock_keluar"`
	Tanggal_keluar                   string                              `json:"tanggal_keluar"`
	Kode_nota                        string                              `json:"kode_nota"`
	Nama_penanggung_jawab            string                              `json:"nama_penanggung_jawab"`
	Nama_toko                        string                              `json:"nama_toko"`
	Jumlah_total                     int                                 `json:"jumlah_total"`
	Total_harga                      int64                               `json:"total_harga"`
	Read_Barang_Stock_Keluar_Request []Read_Barang_Stock_Keluar_Response `json:"read_barang_stock_keluar_request"`
}

type Read_Barang_Stock_Keluar_Response struct {
	Kode_barang_keluar string `json:"kode_barang_keluar"`
	Nama_barang        string `json:"nama_barang"`
	Jumlah_barang      int    `json:"jumlah_barang"`
	Harga_jual         int64  `json:"harga_jual"`
}
