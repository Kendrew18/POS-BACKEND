package request_kasir

type Read_Stock_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
	Kode_store string `json:"kode_store"`
}

type Update_Stock_Kasir_Request struct {
	Harga int64 `json:"harga"`
}

type Update_Stock_Kasir_Kode_Request struct {
	Kode_barang_kasir string `json:"kode_barang_kasir"`
}
