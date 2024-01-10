package request_kasir

type Input_Retur_Customer_Request struct {
	Co                  int    `json:"co"`
	Kode_retur_customer string `json:"kode_retur_customer"`
	Kode_nota           string `json:"kode_nota"`
	Tanggal             string `json:"tanggal"`
	Kode_bentuk_retur   string `json:"kode_bentuk_retur"`
	Kode_store          string `json:"kode_store"`
	Status              int    `json:"status"`
	Kode_kasir          string `json:"kode_kasir"`
}

type Input_Barang_Retur_Customer_Request struct {
	Kode_barang_kasir string `json:"kode_barang_kasir"`
	Jumlah            string `json:"jumlah"`
	Keterangan        string `json:"kode_store"`
}

type Input_Barang_Retur_Customer_Request_V2 struct {
	Co                         int     `json:"co"`
	Kode_barang_retur_customer string  `json:"kode_barang_retur_customer"`
	Kode_barang_kasir          string  `json:"kode_barang_kasir"`
	Kode_retur_customer        string  `json:"kode_retur_customer"`
	Jumlah                     float64 `json:"jumlah"`
	Keterangan                 string  `json:"kode_store"`
}

type Read_Retur_Customer_Request struct {
	Kode_kasir string `json:"kode_kasir"`
}

type Read_Filter_Retur_Customer_Request struct {
	Tanggal    string `json:"tanggal"`
	Kode_store string `json:"kode_store"`
}

type Update_Retur_Customer_Request struct {
	Kode_barang_kasir string  `json:"kode_barang_kasir"`
	Jumlah            float64 `json:"jumlah"`
	Keterangan        string  `json:"keterangan"`
}

type Update_Kode_Retur_Customer_Request struct {
	Kode_barang_retur_customer string `json:"kode_barang_retur_customer"`
}

type Update_Status_Retur_Customer_Request struct {
	Status int `json:"status"`
}

type Update_Status_Retur_Customer_Kode_Request struct {
	Kode_retur_customer string `json:"kode_retur_customer"`
}

type Read_Dropdown_Kode_Nota_Request struct {
	Tanggal    string `json:"tanggal"`
	Kode_kasir string `json:"kode_kasir"`
}

type Read_Dropdown_Barang_Retur_Request struct {
	Kode_nota string `json:"kode_nota"`
}
