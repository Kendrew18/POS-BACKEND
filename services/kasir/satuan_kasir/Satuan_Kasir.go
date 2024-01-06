package satuan_kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	response_kasir "POS-BACKEND/models/response_kasir"
	"net/http"
	"strconv"
)

func Input_Satuan_Kasir(Request request_kasir.Input_Satuan_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("satuan_kasir")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_satuan = "SAT-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_satuan", "nama_satuan", "kode_kasir").Create(&Request)

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

func Read_Satuan_Barang(Request request_kasir.Read_Satuan_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Satuan_Kasir_Response

	con := db.CreateConGorm().Table("satuan_kasir")

	err := con.Select("kode_satuan", "nama_satuan").Where("kode_kasir = ?", Request.Kode_kasir).Order("co ASC").Scan(&data).Error

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

func Delete_Satuan_Barang(Request request_kasir.Delete_Satuan_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	var satuan_barang []string

	con_stock := db.CreateConGorm().Table("barang_kasir")

	err := con_stock.Select("kode_satuan").Where("kode_satuan =?", Request.Kode_satuan).Scan(&satuan_barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if satuan_barang == nil {
		con := db.CreateConGorm().Table("satuan_kasir")

		err := con.Where("kode_satuan=?", Request.Kode_satuan).Delete("")

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
