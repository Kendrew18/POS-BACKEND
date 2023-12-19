package jenis_barang

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Jenis_Barang(jenis_barang request.Input_Jenis_Barang_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("jenis_barang")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	jenis_barang.Co = co + 1
	jenis_barang.Kode_jenis_barang = "JB-" + strconv.Itoa(jenis_barang.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_jenis_barang", "nama_jenis_barang", "kode_gudang").Create(&jenis_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = jenis_barang
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

func Read_Jenis_Barang(kode_gudang request.Read_Jenis_Barang_Request) (response.Response, error) {

	var res response.Response
	var jenis_barang []response.Read_Jenis_Barang_Response

	con := db.CreateConGorm().Table("jenis_barang")

	err := con.Select("kode_jenis_barang", "nama_jenis_barang").Where("kode_gudang=?", kode_gudang.Kode_gudang).Scan(&jenis_barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = jenis_barang
		return res, err
	}

	if jenis_barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = jenis_barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = jenis_barang
	}

	return res, nil
}

func Delete_Jenis_Barang(Request request.Delete_Jenis_Barang_Request) (response.Response, error) {

	var res response.Response

	var jenis_barang []string

	con_masuk := db.CreateConGorm().Table("stock")

	err := con_masuk.Select("kode_jenis_barang").Where("kode_jenis_barang =?", Request.Kode_jenis_barang).Scan(&jenis_barang).Error

	if jenis_barang == nil && err == nil {
		con := db.CreateConGorm().Table("jenis_barang")

		err := con.Where("kode_jenis_barang=?", Request.Kode_jenis_barang).Delete("")

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

func Dropdown_Jenis_Barang(Request request.Dropdown_Jenis_Barang_Request) (response.Response, error) {

	var res response.Response
	var nama_jenis_barang []response.Read_Jenis_Barang_Response

	con := db.CreateConGorm().Table("jenis_barang")

	err := con.Select("kode_jenis_barang", "nama_jenis_barang").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&nama_jenis_barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
	}

	if nama_jenis_barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = nama_jenis_barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = nama_jenis_barang
	}

	return res, nil
}
