package request_kasir

type Input_Pembayaran_Request struct {
	Co                    int     `json:"co"`
	Kode_pembayaran       string  `json:"kode_pembayaran"`
	Kode_nota             string  `json:"kode_nota"`
	Tanggal               string  `json:"tanggal"`
	Kode_jenis_pembayaran string  `json:"kode_jenis_pembayaran"`
	Kode_store            string  `json:"kode_store"`
	Jumlah_total          float64 `json:"jumlah_total"`
	Total_harga           int64   `json:"total_harga"`
	Diskon                int64   `json:"diskon"`
	Kode_kasir            string  `json:"kode_kasir"`
}

type Barang_Input_Pembayaran_Request struct {
	Kode_barang_kasir string `json:"kode_barang_kasir"`
	Jumlah_barang     string `json:"jumlah_barang"`
	Nama_satuan       string `json:"nama_satuan"`
	Harga             string `json:"harga"`
}

type Barang_Input_Pembayaran_Request_V2 struct {
	Co                     int     `json:"co"`
	Kode_pembayaran        string  `json:"kode_pembayaran"`
	Kode_barang_pembayaran string  `json:"kode_barang_pembayaran"`
	Kode_barang_kasir      string  `json:"kode_barang_kasir"`
	Nama_barang_kasir      string  `json:"nama_barang_kasir"`
	Jumlah_barang          float64 `json:"jumlah_barang"`
	Nama_satuan            string  `json:"nama_satuan"`
	Harga                  int64   `json:"harga"`
}

type Update_Jumlah_Dan_Harga_Request struct {
	Jumlah_total float64 `json:"jumlah_total"`
}

type Read_Pembayaran_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}

type Read_Filter_Pembayaran_Request struct {
	Kode_store string `json:"kode_store"`
	Tanggal    string `json:"tanggal"`
}
