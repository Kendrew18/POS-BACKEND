package bentuk_retur

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
	"strconv"
)

func Input_Bentuk_Retur(Request request_kasir.Input_Bentuk_Retur_request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("bentuk_retur")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_satuan = "BR-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_bentuk_retur", "nama_bentuk_retur").Create(&Request)

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

func Read_Bentuk_Retur() (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Satuan_Kasir_Response

	con := db.CreateConGorm().Table("bentuk_retur")

	err := con.Select("kode_bentuk_retur", "nama_bentuk_retur", "nama_satuan").Joins("JOIN satuan_kasir sk ON sk.kode_satuan = bentuk_retur.kode_satuan").Order("co ASC").Scan(&data).Error

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
