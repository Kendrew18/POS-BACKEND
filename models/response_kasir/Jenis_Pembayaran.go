package response_kasir

type Read_Jenis_Pembayaran_Response struct {
	Kode_jenis_pembayaran string `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran string `json:"nama_jenis_pembayaran"`
}
