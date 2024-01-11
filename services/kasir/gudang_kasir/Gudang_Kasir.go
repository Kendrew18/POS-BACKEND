package gudang_kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"net/http"
	"strconv"
)

func Input_Gudang_Kasir(Request request_kasir.Input_Gudang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("gudang_kasir")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_gudang_kasir = "GK-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_gudang_kasir", "kode_kasir", "kode_gudang", "alamat").Create(&Request)

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

func Read_Gudang_Kasir(Request request_kasir.Read_Gudang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Read_Gudang_Kasir_Response

	con := db.CreateConGorm().Table("gudang_kasir")

	err := con.Select("kode_gudang_kasir", "gudang_kasir.kode_gudang", "nama_gudang", "alamat").Joins("JOIN gudang g on g.kode_gudang = gudang_kasir.kode_gudang").Where("gudang_kasir.kode_kasir = ?", Request.Kode_kasir).Order("gudang_kasir.co ASC").Scan(&data).Error

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

func Dropdown_Gudang() (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Dropdown_Gudang_Response

	con := db.CreateConGorm().Table("gudang")

	err := con.Select("kode_gudang", "nama_gudang").Order("gudang.co ASC").Scan(&data).Error

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

func Dropdown_Gudang_Kasir(Request request_kasir.Read_Gudang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data []response_kasir.Dropdown_Gudang_Response

	con := db.CreateConGorm().Table("gudang_kasir")

	err := con.Select("gudang_kasir.kode_gudang AS kode_gudang", "nama_gudang").Joins("JOIN gudang g on g.kode_gudang = gudang_kasir.kode_gudang").Where("gudang_kasir.kode_kasir = ?", Request.Kode_kasir).Order("gudang_kasir.co ASC").Scan(&data).Error

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

func Delete_Gudang_Kasir(Request request_kasir.Delete_Gudang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	var bentuk_retur []string

	con_stock := db.CreateConGorm().Table("barang_request_barang_kasir")

	err := con_stock.Select("kode_gudang_kasir").Where("kode_gudang_kasir = ?", Request.Kode_gudang_kasir).Scan(&bentuk_retur).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if bentuk_retur == nil {
		con := db.CreateConGorm().Table("gudang_kasir")

		err := con.Where("kode_gudang_kasir=?", Request.Kode_gudang_kasir).Delete("")

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
