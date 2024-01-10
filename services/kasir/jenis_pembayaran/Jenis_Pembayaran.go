package jenis_pembayaran

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
	"strconv"
)

func Input_Jenis_Pembayaran(Request request_kasir.Input_Jenis_Pembayaran_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("jenis_pembayaran")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_jenis_pembayaran = "JP-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_jenis_pembayaran", "nama_jenis_pembayaran", "kode_kasir").Create(&Request)

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

func Read_Jenis_Pembayaran(Request request_kasir.Read_Jenis_Pembayaran_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Jenis_Pembayaran_Response

	con := db.CreateConGorm().Table("jenis_pembayaran")

	err := con.Select("kode_jenis_pembayaran", "nama_jenis_pembayaran").Where("kode_kasir = ?", Request.Kode_kasir).Order("co ASC").Scan(&data).Error

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

func Delete_Jenis_Pembayaran(Request request_kasir.Delete_Jenis_Pembayaran_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	var jenis_pembayaran []string

	con_stock := db.CreateConGorm().Table("pembayaran")

	err := con_stock.Select("kode_jenis_pembayaran").Where("kode_jenis_pembayaran =?", Request.Kode_jenis_pembayaran).Scan(&jenis_pembayaran).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if jenis_pembayaran == nil {
		con := db.CreateConGorm().Table("jenis_pembayaran")

		err := con.Where("kode_jenis_pembayaran=?", Request.Kode_jenis_pembayaran).Delete("")

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Erorr karena ada condition yang tidak terpenuhi"
		res.Data = Request
		return res, err
	}

	return res, nil
}
