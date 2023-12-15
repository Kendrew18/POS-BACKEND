package request

type Read_Data_Awal_Audit_Stock_Request struct {
	Kode_gudang      string `json:"kode_gudang"`
	Tanggal_sekarang string `json:"tanggal_sekarang"`
}

type Input_Audit_stock_Request struct {
	Co          int    `json:"co"`
	Kode_audit  string `json:"kode_audit"`
	Tanggal     string `json:"tanggal"`
	Kode_stock  string `json:"kode_stock"`
	Kode_gudang string `json:"kode_gudang"`
}

type Input_Detail_Audit_stock_Request struct {
	Kode_barang_keluar_masuk string `json:"kode_barang_keluar_masuk"`
	Tanggal_masuk            string `json:"tanggal_masuk"`
	Stock_dalam_sistem       string `json:"stock_dalam_sistem"`
	Stock_rill               string `json:"stock_rill"`
	Selisih_stock            string `json:"selisih_stock"`
}

type Input_Detail_Audit_stock_V2_Request struct {
	Co                 int     `json:"co"`
	Kode_detail_audit  string  `json:"kode_detail_audit"`
	Kode_audit         string  `json:"kode_audit"`
	Tanggal_masuk      string  `json:"tanggal_masuk"`
	Stock_dalam_sistem float64 `json:"stock_dalam_sistem"`
	Stock_rill         float64 `json:"stock_rill"`
	Selisih_stock      float64 `json:"selisih_stock"`
	Kode_supplier      string  `json:"kode_supplier"`
}

type Read_Audit_Stock struct {
	Tanggal     string `json:"Tanggal"`
	Kode_gudang string `json:"Kode_gudang"`
}

type Input_Audit_stock_User_Request struct {
	Kode_user string `json:"kode_user"`
}

type Status_Audit_hari_ini_Request struct {
	Status int `json:"status"`
}
