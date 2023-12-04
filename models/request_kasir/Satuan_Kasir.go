package request_kasir

type Input_Satuan_Kasir_Request struct {
	Co          int    `json:"co"`
	Kode_satuan string `json:"kode_satuan"`
	Nama_satuan string `json:"nama_satuan"`
}

type Delete_Satuan_Kasir_Request struct {
	Kode_satuan string `json:"kode_satuan"`
}
