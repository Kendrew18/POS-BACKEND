package barang_kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Barang_Kasir(Request request_kasir.Input_Barang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("barang_kasir")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_barang_kasir = "BK-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_barang_kasir", "nama_barang_kasir", "kode_satuan", "kode_store", "jumlah_pengali", "kode_kasir").Create(&Request)

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

func Read_Barang_Kasir(Request request_kasir.Read_Barang_Kasir_Request) (response_kasir.Response, error) {

	var res response_kasir.Response
	var data response_kasir.Read_Store_Barang_Kasir_Response
	var arr_data []response_kasir.Read_Store_Barang_Kasir_Response

	con_store := db.CreateConGorm().Table("user_management")

	rows, err := con_store.Select("kode_store", "nama_store").Where("kode_kasir=?", Request.Kode_kasir).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		var data_detail []response_kasir.Read_Barang_Kasir_Response

		err_rows := rows.Scan(&data.Kode_store, &data.Nama_store)

		if err_rows != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err_rows
		}

		fmt.Println(data)

		con := db.CreateConGorm().Table("barang_kasir")

		err := con.Select("kode_barang_kasir", "nama_barang_kasir", "nama_satuan", "jumlah_pengali").Joins("JOIN satuan_kasir sk ON sk.kode_satuan = barang_kasir.kode_satuan").Where("barang_kasir.kode_kasir = ? AND kode_store = ?", Request.Kode_kasir, data.Kode_store).Order("barang_kasir.co ASC").Scan(&data_detail).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_barang = data_detail

		arr_data = append(arr_data, data)
	}

	fmt.Println(arr_data)

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}

	return res, nil
}
