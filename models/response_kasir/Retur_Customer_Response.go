package response_kasir

type Read_Retur_Customer_Response struct {
	Kode_retur_customer   string                           `json:"kode_retur_customer"`
	Kode_nota             string                           `json:"kode_nota"`
	Tanggal               string                           `json:"tanggal"`
	Kode_bentuk_retur     string                           `json:"kode_bentuk_retur"`
	Nama_bentuk_retur     string                           `json:"nama_bentuk_retur"`
	Kode_store            string                           `json:"kode_store"`
	Nama_store            string                           `json:"nama_store"`
	Status                int                              `json:"status"`
	Detail_Retur_Customer []Detail_Retur_Customer_Response `json:"detail_retur_customer"`
}

type Detail_Retur_Customer_Response struct {
	Kode_barang_retur_customer string  `json:"kode_barang_retur_customer"`
	Kode_barang_kasir          string  `json:"kode_barang_kasir"`
	Nama_barang_kasir          string  `json:"nama_barang_kasir"`
	Jumlah                     float64 `json:"jumlah"`
	Keterangan                 string  `json:"keterangan"`
}

type Read_Dropdown_Kode_Nota_Request struct {
	Kode_nota  string `json:"kode_nota"`
	Kode_store string `json:"kode_store"`
	Nama_store string `json:"nama_store"`
}
