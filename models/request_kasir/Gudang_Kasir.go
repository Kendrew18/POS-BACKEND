package request_kasir

type Input_Gudang_Kasir_Request struct {
	Co                int    `json:"co"`
	Kode_gudang_kasir string `json:"kode_gudang_kasir"`
	Kode_kasir        string `json:"kode_kasir"`
	Kode_gudang       string `json:"kode_gudang"`
	Alamat            string `json:"alamat"`
}

type Read_Gudang_Kasir_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}
