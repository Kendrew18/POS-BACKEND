package request

type Input_Stock_Keluar_Request struct {
	Co                    int    `json:"co"`
	Kode_stock_keluar     string `json:"kode_stock_keluar"`
	Tanggal_keluar        string `json:"tanggal_keluar"`
	Kode_nota             string `json:"kode_nota"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Kode_toko             string `json:"kode_toko"`
	Kode_gudang           string `json:"kode_gudang"`
}

type Input_Barang_Stock_Keluar_Request struct {
	Kode_stock    string `json:"kode_stock"`
	Jumlah_barang string `json:"jumlah_barang"`
	Harga_jual    string `json:"harga_jual"`
}

type Input_Barang_Stock_Keluar_V2_Request struct {
	Co                       int     `json:"co"`
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Kode_stock_keluar_masuk  string  `json:"kode_stock_keluar_masuk"`
	Kode_stock               string  `json:"kode_stock"`
	Tanggal_kadalurasa       string  `json:"tanggal_kadalurasa"`
	Jumlah_barang            float64 `json:"jumlah_barang"`
	Harga                    int64   `json:"harga"`
	Total_harga              int64   `json:"total_harga"`
	Status                   int     `json:"status"`
}

type Read_Stock_Keluar_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Filter_Stock_Keluar_Request struct {
	Tanggal_1 string `json:"tanggal_1"`
	Tanggal_2 string `json:"tanggal_2"`
	Kode_toko string `json:"kode_toko"`
}
