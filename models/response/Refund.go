package response

type Read_Refund_Response struct {
	Kode_refund          string                        `json:"kode_refund"`
	Tanggal              string                        `json:"tanggal"`
	Tanggal_pengambalian string                        `json:"tanggal_pengembalian"`
	Nama_supplier        string                        `json:"nama_supplier"`
	Jumlah_total         float64                       `json:"jumlah_total"`
	Status               int                           `json:"status"`
	Detail_stock_masuk   []Read_Barang_Refund_Response `json:"detail_stock_masuk"`
}

type Read_Barang_Refund_Response struct {
	Kode_nota           string  `json:"kode_nota"`
	Nama_barang         string  `json:"nama_barang"`
	Tanggal_stock_masuk string  `json:"tanggal_stock_masuk"`
	Jumlah              float64 `json:"jumlah"`
	Keterangan          string  `json:"keterangan"`
}
