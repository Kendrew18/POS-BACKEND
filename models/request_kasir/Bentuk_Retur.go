package request_kasir

type Input_Bentuk_Retur_Request struct {
	Co                int    `json:"co"`
	Kode_bentuk_retur string `json:"kode_bentuk_retur"`
	Nama_bentuk_retur string `json:"nama_bentuk_retur"`
	Kode_kasir        string `json:"kode_kasir"`
}

type Read_Bentuk_Retur_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}
