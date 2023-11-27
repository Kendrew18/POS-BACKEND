package audit

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"net/http"
)

func Read_Data_Awal_Audit_Stock(Request request.Read_Data_Awal_Audit_Stock_Request) (response.Response, error) {
	var res response.Response
	var data response.Read_Awal_Audit_Stock_Response
	var arr_data []response.Read_Awal_Audit_Stock_Response

	con := db.CreateConGorm().Table("stock")

	rows, err := con.Select("kode_stock", "nama_barang", "jumlah").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		var detail_data []response.Detail_Aduit_Awal_Stock_Response
		err = rows.Scan(&data.Kode_stock, &data.Nama_Barang, &data.Jumlah)
		data.Tanggal_Sekarang = Request.Tanggal_sekarang

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}
		con_detail_barang := db.CreateConGorm().Table("detail_stock")

		err2 := con_detail_barang.Select("kode_barang_keluar_masuk", "tanggal", "jumlah_barang").Joins("JOIN stock_keluar_masuk skm on skm.kode_stock_keluar_masuk = detail_stock.kode_stock_keluar_masuk").Where("kode_stock =?", data.Kode_stock).Scan(&detail_data)

		if err2.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err2.Error
		}

		data.Detail_audit_awal = detail_data

		arr_data = append(arr_data, data)
	}

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

func Input_Audit_Stock(Request request.Read_Data_Awal_Audit_Stock_Request) (response.Response, error) {
	var res response.Response
	var data response.Read_Awal_Audit_Stock_Response
	var arr_data []response.Read_Awal_Audit_Stock_Response

	con := db.CreateConGorm().Table("stock")

	rows, err := con.Select("kode_stock", "nama_barang", "jumlah").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		var detail_data []response.Detail_Aduit_Awal_Stock_Response
		err = rows.Scan(&data.Kode_stock, &data.Nama_Barang, &data.Jumlah)
		data.Tanggal_Sekarang = Request.Tanggal_sekarang

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}
		con_detail_barang := db.CreateConGorm().Table("detail_stock")

		err2 := con_detail_barang.Select("kode_barang_keluar_masuk", "tanggal", "jumlah_barang").Joins("JOIN stock_keluar_masuk skm on skm.kode_stock_keluar_masuk = detail_stock.kode_stock_keluar_masuk").Where("kode_stock =?", data.Kode_stock).Scan(&detail_data)

		if err2.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err2.Error
		}

		data.Detail_audit_awal = detail_data

		arr_data = append(arr_data, data)
	}

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
