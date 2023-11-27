package response

type Read_Pre_Order_Response struct {
	Kode_pre_order     string                             `json:"kode_pre_order"`
	Tanggal            string                             `json:"tanggal"`
	Kode_nota          string                             `json:"kode_nota"`
	Penanggung_jawab   string                             `json:"penanggung_jawab"`
	Nama_supplier      string                             `json:"nama_supplier"`
	Jumlah_total       float64                            `json:"jumlah_total"`
	Total_harga        int64                              `json:"total_harga"`
	Status             int                                `json:"status"`
	Detail_stock_masuk []Read_Detail_Stock_Masuk_Response `json:"detail_stock_masuk"`
}

type Read_Barang_Pre_Order_Response struct {
	Kode_barang_pre_order string  `json:"kode_barang_pre_order"`
	Nama_barang           string  `json:"nama_barang"`
	Tanggal_kadaluarsa    string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang         float64 `json:"jumlah_barang"`
	Harga                 int64   `json:"harga"`
}
