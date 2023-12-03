package request

type Read_Kartu_Stock_Request struct {
	Kode_gudang   string `json:"kode_gudang"`
	Tanggal_1     string `json:"tanggal_1"`
	Tanggal_2     string `json:"tanggal_2"`
	Kode_supplier string `json:"kode_supplier"`
	Kode_stock    string `json:"kode_stock"`
}
