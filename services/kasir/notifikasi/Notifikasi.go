package notifikasi

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
)

func Read_Notifikasi(Request request_kasir.Read_Notifikasi_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Notifikasi_Kasir_Response

	con := db.CreateConGorm().Table("barang_kasir")

	err := con.Select("kode_barang_kasir", "nama_barang_kasir", "jumlah", "jumlah_minimal", "IF(jumlah >= jumlah_minimal,0,1) as status").Joins("JOIN satuan_kasir sk ON sk.kode_satuan = barang_kasir.kode_satuan").Where("barang_kasir.kode_kasir = ?", Request.Kode_kasir).Order("barang_kasir.co ASC").Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
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

func Update_Jumlah_Minimal(Request request_kasir.Update_Jumlah_Minimal_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	con := db.CreateConGorm().Table("barang_kasir")

	err := con.Where("kode_barang_kasir = ?", Request.Kode_barang_kasir).Update("jumlah_minimal", Request.Jumlah_minimal)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}
	return res, nil
}
