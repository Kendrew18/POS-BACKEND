package request

type Read_Kartu_Stock_Request struct {
	Kode_gudang string `json:"kode_gudang"`
	Tanggal     string `json:"tanggal"`
}
