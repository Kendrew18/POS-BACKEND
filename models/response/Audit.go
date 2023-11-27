package response

type Read_Awal_Audit_Stock_Response struct {
	Kode_stock        string                             `json:"kode_stock"`
	Tanggal_Sekarang  string                             `json:"tanggal_Sekarang"`
	Nama_Barang       string                             `json:"nama_Barang"`
	Jumlah            float64                            `json:"jumlah"`
	Detail_audit_awal []Detail_Aduit_Awal_Stock_Response `json:"detail_audit_awal"`
}

type Detail_Aduit_Awal_Stock_Response struct {
	Kode_barang_keluar_masuk string `json:"kode_stock_keluar_masuk"`
	Tanggal                  string `json:"tanggal"`
	Jumlah_barang            string `json:"jumlah_barang"`
}
