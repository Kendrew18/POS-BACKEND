package response

type Read_Awal_Audit_Stock_Response struct {
	Kode_stock        string                             `json:"kode_stock"`
	Tanggal_Sekarang  string                             `json:"tanggal_Sekarang"`
	Nama_Barang       string                             `json:"nama_Barang"`
	Jumlah            float64                            `json:"jumlah"`
	Detail_audit_awal []Detail_Aduit_Awal_Stock_Response `json:"detail_audit_awal"`
}

type Detail_Aduit_Awal_Stock_Response struct {
	Kode_barang_keluar_masuk string  `json:"kode_stock_keluar_masuk"`
	Tanggal                  string  `json:"tanggal"`
	Jumlah_barang            float64 `json:"jumlah_barang"`
}

type Read_Audit_Stock_Response struct {
	Kode_audit                 string                        `json:"kode_audit"`
	Tanggal                    string                        `json:"tanggal"`
	Nama_barang                string                        `json:"nama_barang"`
	Total_jumlah_dalam_sistem  float64                       `json:"total_jumlah_dalam_sistem"`
	Total_jumlah_stock_rill    float64                       `json:"total_jumlah_stock_rill"`
	Total_jumlah_selisih_stock float64                       `json:"total_jumlah_selisih_stock"`
	Detail_audit_awal          []Detail_Aduit_Stock_Response `json:"detail_audit_awal"`
}

type Detail_Aduit_Stock_Response struct {
	Kode_detail_audit  string  `json:"kode_detail_audit"`
	Tanggal_masuk      string  `json:"tanggal_masuk"`
	Stock_dalam_sistem float64 `json:"stock_dalam_sistem"`
	Stock_rill         float64 `json:"stock_rill"`
	Selisih_stock      float64 `json:"selisih_stock"`
}
