package response

type Read_Stock_Masuk_Response struct {
	Kode_stock_masuk   string                             `json:"kode_stock_masuk"`
	Tanggal_masuk      string                             `json:"tanggal_masuk"`
	Kode_nota          string                             `json:"kode_nota"`
	Penanggung_jawab   string                             `json:"penanggung_jawab"`
	Nama_supplier      string                             `json:"nama_supplier"`
	Jumlah_total       int                                `json:"jumlah"`
	Total_harga        int64                              `json:"total_harga"`
	Detail_stock_masuk []Read_Detail_Stock_Masuk_Response `json:"detail_stock_masuk"`
}

type Read_Detail_Stock_Masuk_Response struct {
	Kode_barang_masuk  string  `json:"kode_barang_masuk"`
	Nama_barang        string  `json:"nama_barang"`
	Tanggal_kadaluarsa string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang      float64 `json:"jumlah_barang"`
	Harga_pokok        int64   `json:"harga_pokok"`
}
