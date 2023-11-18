package response

type Read_Stock_Response struct {
	Kode_stock         string                  `json:"kode_stock"`
	Nama_Barang        string                  `json:"nama_barang"`
	Harga_jual         int64                   `json:"harga_jual"`
	Jumlah             int                     `json:"jumlah"`
	Nama_satuan_barang string                  `json:"nama_satuan_barang"`
	Nama_jenis_barang  string                  `json:"nama_jenis_barang"`
	Detail_stock       []Detail_Stock_Response `json:"detail_stock"`
}

type Detail_Stock_Response struct {
	Tanggal_masuk string  `json:"tanggal_masuk"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Harga         int64   `json:"harga"`
}

type Read_Barang_Response struct {
	Kode_jenis_barang  string                        `json:"kode_jenis_barang"`
	Jenis_Barang       string                        `json:"jenis_barang"`
	Read_Detail_Barang []Read_Detail_Barang_Response `json:"read_detail_barang"`
}

type Read_Detail_Barang_Response struct {
	Kode_stock         string `json:"kode_stock"`
	Nama_Barang        string `json:"nama_barang"`
	Harga_jual         int64  `json:"harga_jual"`
	Nama_satuan_barang string `json:"nama_satuan_barang"`
}

type Dropdown_Nama_Barang_Response struct {
	Kode_stock  string `json:"kode_stock"`
	Nama_barang string `json:"nama_barang"`
}

type Read_Detail_Stock_Response struct {
	Kode_barang_keluar_masuk string `json:"kode_barang_keluar_masuk"`
}
