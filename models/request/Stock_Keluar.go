package request

type Input_Stock_Keluar_Request struct {
	Co                    int    `json:"co"`
	Kode_stock_keluar     string `json:"kode_stock_keluar"`
	Kode_gudang           string `json:"kode_gudang"`
	Tanggal_keluar        string `json:"tanggal_keluar"`
	Kode_nota             string `json:"kode_nota"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Kode_toko             string `json:"kode_toko"`
}

type Input_Barang_Stock_Keluar_Request struct {
	Kode_stock    string `json:"kode_stock"`
	Jumlah_barang string `json:"jumlah_barang"`
	Harga_jual    string `json:"harga_jual"`
}

type Input_Barang_Stock_Keluar_V2_Request struct {
	Co                 int     `json:"co"`
	Kode_barang_keluar string  `json:"kode_barang_keluar"`
	Kode_stock_keluar  string  `json:"kode_stock_keluar"`
	Kode_stock         string  `json:"kode_stock"`
	Tanggal_kadalurasa string  `json:"tanggal_kadalurasa"`
	Jumlah_barang      float64 `json:"jumlah_barang"`
	Harga_jual         int64   `json:"harga_jual"`
	Total_harga        int64   `json:"total_harga"`
}

type Read_Stock_Keluar_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
