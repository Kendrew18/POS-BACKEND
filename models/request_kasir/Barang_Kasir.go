package request_kasir

type Input_Barang_Kasir_Request struct {
	Co                int    `json:"co"`
	Kode_barang_kasir string `json:"kode_barang_kasir"`
	Nama_barang_kasir string `json:"nama_barang_kasir"`
	Kode_satuan       string `json:"kode_satuan"`
	Kode_store        string `json:"kode_store"`
	Jumlah_pengali    string `json:"jumlah_pengali"`
	Kode_kasir        string `json:"kode_kasir"`
}

type Read_Barang_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}

type Dropdown_Barang_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
	Kode_store string `json:"kode_store"`
}
