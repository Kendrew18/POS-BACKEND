package request

type Read_Data_Awal_Audit_Stock_Request struct {
	Kode_gudang      string `json:"kode_gudang"`
	Tanggal_sekarang string `json:"tanggal_sekarang"`
}
