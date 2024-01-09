package stock_kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
)

func Read_Stock_Kasir(Request request_kasir.Read_Stock_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Stock_Kasir_Response

	con := db.CreateConGorm().Table("barang_kasir")

	err := con.Select("kode_barang_kasir", "nama_barang_kasir", "nama_satuan", "jumlah", "harga").Joins("JOIN satuan_kasir sk ON sk.kode_satuan = barang_kasir.kode_satuan").Where("barang_kasir.kode_kasir = ? AND kode_store = ?", Request.Kode_kasir, Request.Kode_store).Order("barang_kasir.co ASC").Scan(&data).Error

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

func Update_Harga_Stock_Kasir(Request request_kasir.Update_Stock_Kasir_Request, Request_kode request_kasir.Update_Stock_Kasir_Kode_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	con := db.CreateConGorm()

	err := con.Table("barang_kasir").Where("kode_barang_kasir = ?", Request_kode.Kode_barang_kasir).Select("harga").Updates(&Request)

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
