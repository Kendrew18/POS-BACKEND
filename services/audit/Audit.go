package audit

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

func Input_Audit_Stock(Request request.Input_Audit_stock_Request, Request_detail request.Input_Detail_Audit_stock_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("audit")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_audit = "AU-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	err = con.Select("co", "kode_audit", "tanggal", "kode_stock", "kode_gudang").Create(&Request)

	tanggal_masuk := tools.String_Separator_To_String(Request_detail.Tanggal_masuk)
	stock_dalam_sistem := tools.String_Separator_To_float64(Request_detail.Stock_dalam_sistem)
	stock_rill := tools.String_Separator_To_float64(Request_detail.Stock_rill)
	selisih_stock := tools.String_Separator_To_float64(Request_detail.Selisih_stock)

	for i := 0; i < len(tanggal_masuk); i++ {
		var detail request.Input_Detail_Audit_stock_V2_Request

		con_detail := db.CreateConGorm().Table("detail_audit")

		co := 0

		err := con_detail.Select("co").Order("co DESC").Limit(1).Scan(&co)

		detail.Co = co + 1
		detail.Kode_detail_audit = "DAU-" + strconv.Itoa(detail.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = detail
			return res, err.Error
		}

		detail.Kode_audit = Request.Kode_audit
		detail.Stock_rill = stock_rill[i]
		detail.Stock_dalam_sistem = stock_dalam_sistem[i]
		detail.Selisih_stock = selisih_stock[i]

		date3, _ := time.Parse("02-01-2006", tanggal_masuk[i])
		detail.Tanggal_masuk = date3.Format("2006-01-02")

		fmt.Println(detail.Tanggal_masuk)
		fmt.Println(detail)

		err = con_detail.Select("co", "kode_detail_audit", "kode_audit", "tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock").Create(&detail)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = detail
			return res, err.Error
		}
	}

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

func Read_Audit_Stock(Request request.Read_Audit_Stock) (response.Response, error) {
	var res response.Response
	var data response.Read_Audit_Stock_Response
	var arr_data []response.Read_Audit_Stock_Response

	con := db.CreateConGorm().Table("audit")

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	rows, err := con.Select("kode_audit", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "nama_barang", "SUM(ds.stock_dalam_sistem)", "SUM(stock_rill)", "SUM(selisih_stock)").Joins("JOIN stock s ON s.kode_stock = audit.kode_stock").Joins("JOIN detail_stock ds ON ds.kode_audit = audit.kode_audit").Where("kode_gudang = ? && tanggal = ?", Request.Kode_gudang, Request.Tanggal).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		var detail_data []response.Detail_Aduit_Stock_Response
		err = rows.Scan(&data.Kode_audit, &data.Tanggal, &data.Nama_barang, &data.Total_jumlah_dalam_sistem, &data.Total_jumlah_stock_rill, &data.Total_jumlah_selisih_stock)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}
		con_detail_audit := db.CreateConGorm().Table("detail_audit")

		err2 := con_detail_audit.Select("kode_detail_audit", "DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock").Where("kode_audit = ?", data.Kode_audit).Scan(&detail_data)

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
