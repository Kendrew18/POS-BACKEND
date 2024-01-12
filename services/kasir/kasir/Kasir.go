package kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
)

func Read_Menu_Kasir(Request request_kasir.Read_Stock_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Stock_Kasir_Response

	con := db.CreateConGorm().Table("barang_kasir")

	err := con.Select("kode_barang_kasir", "nama_barang_kasir", "nama_satuan", "jumlah*100/100 as jumlah", "harga").Joins("JOIN satuan_kasir sk ON sk.kode_satuan = barang_kasir.kode_satuan").Where("barang_kasir.kode_kasir = ? AND kode_store = ? AND jumlah > 0", Request.Kode_kasir, Request.Kode_store).Order("barang_kasir.co ASC").Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}

	return res, nil
}
