package response_kasir

type Read_Gudang_Kasir_Response struct {
	Kode_gudang_kasir string `json:"kode_gudang_kasir"`
	Kode_gudang       string `json:"kode_gudang"`
	Nama_gudang       string `json:"nama_gudang"`
	Alamat            string `json:"alamat"`
}
