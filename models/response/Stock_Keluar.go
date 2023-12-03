package response

type Read_Stock_Keluar_Response struct {
	Kode_stock_keluar_masuk          string                              `json:"kode_stock_keluar_masuk"`
	Tanggal                          string                              `json:"tanggal"`
	Kode_nota                        string                              `json:"kode_nota"`
	Penanggung_jawab                 string                              `json:"penanggung_jawab"`
	Nama_toko                        string                              `json:"nama_toko"`
	Jumlah_total                     float64                             `json:"jumlah_total"`
	Total_harga                      int64                               `json:"total_harga"`
	Read_Barang_Stock_Keluar_Request []Read_Barang_Stock_Keluar_Response `json:"read_barang_stock_keluar_request"`
}

type Read_Barang_Stock_Keluar_Response struct {
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Nama_barang              string  `json:"nama_barang"`
	Jumlah_barang            float64 `json:"jumlah_barang"`
	Harga                    int64   `json:"harga_jual"`
}

type Read_Tanggal_dan_Jumlah struct {
	Kode_stock_keluar_masuk  string  `json:"kode_stock_keluar_masuk"`
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Kode_stock               string  `json:"kode_stock"`
	Tanggal_masuk            string  `json:"Tanggal_masuk"`
	Kode                     string  `json:"kode"`
	Jumlah_barang            float64 `json:"Jumlah_barang"`
}

type Kode_stock_keluar_masuk struct {
	Co                       int    `json:"co"`
	Kode_pengurangan         string `json:"kode_pengurangan"`
	Kode_stock_keluar_masuk  string `json:"Kode_stock_keluar_masuk"`
	Kode_barang_keluar_masuk string `json:"Kode_barang_keluar_masuk"`
	Kode_stock_keluar        string `json:"Kode_stock_keluar"`
	Kode_barang_keluar       string `json:"kode_barang_keluar"`
	Kode_supplier            string `json:"kode_supplier"`
}
