package toko

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Toko(Request request.Input_Toko_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("toko")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_toko = "TK-" + strconv.Itoa(Request.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_toko", "nama_toko", "alamat", "nomor_telpon", "kode_gudang").Create(&Request)

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

func Read_Toko(Request request.Read_Toko_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Toko_Response

	con := db.CreateConGorm().Table("toko")

	err := con.Select("kode_toko", "nama_toko", "alamat", "nomor_telpon").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&data).Error

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

func Delete_Toko(Request request.Delete_Toko_Request) (response.Response, error) {
	var res response.Response

	var barang_stock_keluar []string

	con_keluar := db.CreateConGorm().Table("stock_keluar")

	err := con_keluar.Select("kode_toko").Where("kode_toko =?", Request.Kode_toko).Scan(&barang_stock_keluar).Error

	if barang_stock_keluar == nil && err == nil {
		con := db.CreateConGorm().Table("toko")

		err := con.Where("kode_toko=?", Request.Kode_toko).Delete("")

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

func Dropdown_Nama_Toko(Request request.Read_Toko_Request) (response.Response, error) {
	var res response.Response
	var data []response.Read_Dropdown_Nama_Toko_Response

	con := db.CreateConGorm().Table("toko")

	err := con.Select("kode_toko", "nama_toko").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&data).Error

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
