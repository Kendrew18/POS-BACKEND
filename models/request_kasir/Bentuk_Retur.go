package request_kasir

type Input_Bentuk_Retur_request struct {
	Co                int    `json:"co"`
	Kode_bentuk_retur string `json:"kode_bentuk_retur"`
	Nama_bentuk_retur string `json:"nama_bentuk_retur"`
	Kode_satuan       string `json:"kode_satuan"`
}
