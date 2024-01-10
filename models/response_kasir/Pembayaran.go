package response_kasir

type Read_Pembayaran_Response struct {
	Kode_pembayaran       string                            `json:"kode_pembayaran"`
	Kode_nota             string                            `json:"kode_nota"`
	Tanggal               string                            `json:"tanggal"`
	Kode_jenis_pembayaran string                            `json:"kode_jenis_pembayaran"`
	Nama_jenis_pambayaran string                            `json:"nama_jenis_pambayaran"`
	Nama_store            string                            `json:"nama_store"`
	Jumlah_total          float64                           `json:"jumlah_total"`
	Total_harga           int64                             `json:"total_harga"`
	Detail_pembayaran     []Read_Detail_Pembayaran_Response `json:"detail_pembayaran"`
}

type Read_Detail_Pembayaran_Response struct {
	Kode_barang_pembayaran string  `json:"kode_barang_pembayaran"`
	Kode_barang_kasir      string  `json:"kode_barang_kasir"`
	Nama_barang_kasir      string  `json:"nama_barang_kasir"`
	Jumlah_barang          float64 `json:"jumlah_barang"`
	Nama_satuan            string  `json:"nama_satuan"`
	Harga                  int64   `json:"harga"`
}
