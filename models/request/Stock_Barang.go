package request

type Input_Barang_Request struct {
	Nama_Barang        string `json:"nama_barang"`
	Harga_jual         string `json:"harga_jual"`
	Kode_satuan_barang string `json:"kode_satuan_barang"`
	Kode_jenis_barang  string `json:"kode_jenis_barang"`
	Kode_gudang        string `json:"kode_gudang"`
}

type Input_Barang_Request_V2 struct {
	Co                 int    `json:"co"`
	Kode_stock         string `json:"kode_stock"`
	Nama_Barang        string `json:"nama_barang"`
	Harga_jual         int64  `json:"harga_jual"`
	Kode_satuan_barang string `json:"kode_satuan_barang"`
	Kode_jenis_barang  string `json:"kode_jenis_barang"`
	Kode_gudang        string `json:"kode_gudang"`
}

type Read_Stock_Request struct {
	Kode_gudang       string `json:"kode_gudang"`
	Kode_jenis_barang string `json:"kode_jenis_barang"`
}

type Delete_Barang_Request struct {
	Kode_stock string `json:"kode_stock"`
}

type Dropdown_Nama_Barang_Request struct {
	Kode_supplier string `json:"Kode_supplier"`
}

type Read_Detail_Stock struct {
	Kode_stock  string `json:"kode_stock"`
	Kode_gudang string `json:"kode_gudang"`
}
