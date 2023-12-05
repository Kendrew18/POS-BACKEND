package request_kasir

type Input_Satuan_Kasir_Request struct {
	Co          int    `json:"co"`
	Kode_satuan string `json:"kode_satuan"`
	Nama_satuan string `json:"nama_satuan"`
	Kode_kasir  string `json:"kode_kasir"`
}

type Read_Satuan_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}

type Delete_Satuan_Kasir_Request struct {
	Kode_satuan string `json:"kode_satuan"`
}
