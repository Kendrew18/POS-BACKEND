package response

type Read_Supplier_Response struct {
	Kode_supplier   string                          `json:"kode_supplier"`
	Nama_supplier   string                          `json:"nama_supplier"`
	Nomor_telpon    string                          `json:"nomor_telpon"`
	Barang_supplier []Read_Barang_Supplier_Response `gorm:"many2many:Read_Barang_Supplier_Response" json:"barang_supplier"`
}

type Read_Barang_Supplier_Response struct {
	Kode_stock  string `json:"kode_stock"`
	Nama_barang string `json:"nama_barang"`
}

type Read_Dropdown_Nama_Supplier_Response struct {
	Kode_supplier string `json:"kode_supplier"`
	Nama_supplier string `json:"nama_supplier"`
}
