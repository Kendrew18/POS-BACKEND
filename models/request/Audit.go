package request

type Read_Data_Awal_Audit_Stock_Request struct {
	Kode_gudang      string `json:"kode_gudang"`
	Tanggal_sekarang string `json:"tanggal_sekarang"`
}

type Input_Audit_stock_Request struct {
	Co          int    `json:"co"`
	Kode_audit  string `json:"kode_audit"`
	Kode_stock  string `json:"kode_stock"`
	Kode_gudang string `json:"kode_gudang"`
	Status      int    `json:"status"`
}

type Input_Detail_Audit_stock_Request struct {
	Co                       int     `json:"co"`
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Kode_detail_audit        string  `json:"kode_detail_audit"`
	Kode_audit               string  `json:"kode_audit"`
	Tanggal_masuk            string  `json:"tanggal_masuk"`
	Stock_dalam_sistem       float64 `json:"stock_dalam_sistem"`
	Stock_rill               float64 `json:"stock_rill"`
	Selisih_stock            float64 `json:"selisih_stock"`
	Kode_supplier            string  `json:"kode_supplier"`
	Status                   int     `json:"status"`
}

type Update_Stock_Rill struct {
	Stock_rill         float64 `json:"stock_rill"`
	Selisih_stock      float64 `json:"selisih_stock"`
	Stock_dalam_sistem float64 `json:"stock_dalam_sistem"`
}

type Update_Status_Audit_Request struct {
	Kode_audit  string `json:"kode_audit"`
	Tanggal     string `json:"tanggal"`
	Kode_user   string `json:"kode_user"`
	Kode_gudang string `json:"kode_gudang"`
}

type Update_Status_Audit_Request_V2 struct {
	Status  int    `json:"status"`
	Tanggal string `json:"tanggal"`
}

type Read_Audit_Stock struct {
	Tanggal     string `json:"Tanggal"`
	Kode_gudang string `json:"Kode_gudang"`
}

type Status_Audit_hari_ini_Request struct {
	Status int `json:"status"`
}
