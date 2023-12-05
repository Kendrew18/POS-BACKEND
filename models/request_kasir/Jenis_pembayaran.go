package request_kasir

type Input_Jenis_Pembayaran_Request struct {
	Co                    int    `json:"co"`
	Kode_jenis_pembayaran string `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran string `json:"nama_jenis_pembayaran"`
	Kode_kasir            string `json:"kode_kasir"`
}

type Read_Jenis_Pembayaran_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}
