package response

type Read_Sup_Plus_Barang_Response struct {
	Kode_stock    string `json:"kode_stock"`
	Nama_barang   string `json:"nama_barang"`
	Kode_supplier string `json:"kode_supplier"`
	Nama_supplier string `json:"nama_supplier"`
}

type Read_Raw_Kartu_Stock_Response struct {
	Kode_stock_keluar_masuk string `json:"kode_stock_keluar_masuk"`
	Tanggal                 string `json:"tanggal"`
	Kode                    string `json:"kode"`
}

type Kartu_Stock_Response struct {
	Nama_barang         string                        `json:"nama_barang"`
	Nama_supplier       string                        `json:"nama_supplier"`
	Jumlah_stock_masuk  float64                       `json:"jumlah_stock_masuk"`
	Jumlah_stock_keluar float64                       `json:"jumlah_stock_keluar"`
	Detail_kartu_stock  []Detail_Kartu_Stock_Response `json:"detail_kartu_stock"`
}

type Detail_Kartu_Stock_Response struct {
	Tanggal       string  `json:"tanggal"`
	Keterangan    string  `json:"keterangan"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Sisa          float64 `json:"sisa"`
}
