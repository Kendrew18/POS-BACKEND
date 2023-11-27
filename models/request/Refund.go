package request

type Input_Refund_Request struct {
	Co                   int    `json:"co"`
	Kode_refund          string `json:"kode_refund"`
	Tanggal              string `json:"tanggal"`
	Tanggal_pengembalian string `json:"tanggal_pengembalian"`
	Kode_supplier        string `json:"kode_supplier"`
	Kode_gudang          string `json:"kode_gudang"`
}

type Input_Barang_Refund_Request struct {
	Kode_nota           string `json:"kode_nota"`
	Kode_stock          string `json:"kode_stock"`
	Tanggal_stock_masuk string `json:"Tanggal_stock_masuk"`
	Jumlah              string `json:"jumlah"`
	Keterangan          string `json:"keterangan"`
}

type Input_Barang_Refund_V2_Request struct {
	Co                  int     `json:"co"`
	Kode_barang_refund  string  `json:"kode_barang_refund"`
	Kode_nota           string  `json:"kode_nota"`
	Kode_stock          string  `json:"kode_stock"`
	Tanggal_stock_masuk string  `json:"Tanggal_stock_masuk"`
	Jumlah              float64 `json:"jumlah"`
	Keterangan          string  `json:"keterangan"`
	Kode_refund         string  `json:"kode_refund"`
}

type Read_Refund_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Refund_Filter_Request struct {
	Tanggal_1     string `json:"tanggal_1"`
	Tanggal_2     string `json:"tanggal_2"`
	Kode_supplier string `json:"kode_supplier"`
}

type Update_Refund_Request struct {
	Kode_barang_refund string `json:"kode_barang_refund"`
}

type Update_Barang_Refund_Request struct {
	Kode_nota           string  `json:"kode_nota"`
	Kode_stock          string  `json:"kode_stock"`
	Tanggal_stock_masuk string  `json:"Tanggal_stock_masuk"`
	Jumlah              float64 `json:"Jumlah"`
	Keterangan          string  `json:"Keterangan"`
}

type Update_Status_Refund_Request struct {
	Status int `json:"Status"`
}

type Update_Status_Kode_Refund_Request struct {
	Kode_refund string `json:"kode_refund"`
}

type Move_Refund_To_Stock_Keluar_Request struct {
	Kode_stock    string  `json:"kode_stock"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Harga_jual    int64   `json:"harga_jual"`
}

type Tanggal_dan_Kode_Gudang struct {
	Tanggal_pengembalian string `json:"tanggal_pengembalian"`
	Kode_gudang          string `json:"kode_gudang"`
}

type Data_Update_Refund struct {
	Tanggal_pengembalian string `json:"tanggal_pengembalian"`
	Status               int    `json:"status"`
}
