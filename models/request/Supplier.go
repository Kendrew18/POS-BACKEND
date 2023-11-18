package request

type Input_Supplier_Request struct {
	Co            int    `json:"co"`
	Kode_supplier string `json:"kode_supplier"`
	Nama_supplier string `json:"nama_supplier"`
	Nomor_telpon  string `json:"nomor_telpon"`
	Kode_gudang   string `json:"kode_gudang"`
}

type Input_Barang_Supplier_Request struct {
	Co                   int    `json:"co"`
	Kode_barang_supplier string `json:"kode_barang_supplier"`
	Kode_supplier        string `json:"kode_supplier"`
	Kode_stock           string `json:"kode_stock"`
}

type Read_Supplier_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Delete_Supplier_Request struct {
	Kode_supplier string `json:"kode_supplier"`
}
