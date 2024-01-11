package bentuk_retur

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
	"strconv"
)

func Input_Bentuk_Retur(Request request_kasir.Input_Bentuk_Retur_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("bentuk_retur")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_bentuk_retur = "BR-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_bentuk_retur", "nama_bentuk_retur", "kode_kasir").Create(&Request)

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

func Read_Bentuk_Retur(Request request_kasir.Read_Barang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Bentuk_Retur_Response

	con := db.CreateConGorm().Table("bentuk_retur")

	err := con.Select("kode_bentuk_retur", "nama_bentuk_retur").Where("kode_kasir = ?", Request.Kode_kasir).Order("co ASC").Scan(&data).Error

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

func Delete_Bentuk_Retur(Request request_kasir.Delete_Bentuk_Retur_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	var bentuk_retur []string

	con_stock := db.CreateConGorm().Table("retur_customer")

	err := con_stock.Select("kode_bentuk_retur").Where("kode_bentuk_retur =?", Request.Kode_bentuk_retur).Scan(&bentuk_retur).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if bentuk_retur == nil || (Request.Kode_bentuk_retur != "BR-1" && Request.Kode_bentuk_retur != "BR-2") {
		con := db.CreateConGorm().Table("bentuk_retur")

		err := con.Where("kode_bentuk_retur=?", Request.Kode_bentuk_retur).Delete("")

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
