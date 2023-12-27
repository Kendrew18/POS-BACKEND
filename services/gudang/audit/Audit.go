package audit

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"fmt"
	"math"
	"strconv"

	"net/http"
	"time"

	"gorm.io/gorm"
)

func Input_Audit_Stock(Request request.Input_Audit_stock_Request, Request_detail request.Input_Detail_Audit_stock_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	if Request.Kode_audit == "" && Request_detail.Kode_detail_audit == "" {

		co := 0

		err := con.Table("audit").Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Kode_audit = "AU-" + strconv.Itoa(Request.Co)
		Request.Status = 0

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("audit").Select("co", "kode_audit", "kode_stock", "kode_gudang", "status").Create(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		co = 0
		err = con.Table("detail_audit").Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_detail.Co = co + 1
		Request_detail.Kode_detail_audit = "DAU-" + strconv.Itoa(Request_detail.Co)
		Request_detail.Kode_audit = Request.Kode_audit
		Request_detail.Status = 0

		date, _ := time.Parse("02-01-2006", Request_detail.Tanggal_masuk)
		Request_detail.Tanggal_masuk = date.Format("2006-01-02")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("detail_audit").Select("co", "kode_barang_keluar_masuk", "kode_detail_audit", "kode_audit", "tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock", "kode_supplier", "status").Create(&Request_detail)

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

	} else if Request.Kode_audit != "" && Request_detail.Kode_detail_audit == "" {

		co := 0
		err := con.Table("detail_audit").Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_detail.Co = co + 1
		Request_detail.Kode_detail_audit = "DAU-" + strconv.Itoa(Request_detail.Co)
		Request_detail.Kode_audit = Request.Kode_audit
		Request_detail.Status = 0

		date, _ := time.Parse("02-01-2006", Request_detail.Tanggal_masuk)
		Request_detail.Tanggal_masuk = date.Format("2006-01-02")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("detail_audit").Select("co", "kode_barang_keluar_masuk", "kode_detail_audit", "kode_audit", "tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock", "kode_supplier", "status").Create(&Request_detail)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_detail
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}

	} else if Request.Kode_audit != "" && Request_detail.Kode_detail_audit != "" {

		var stock_rill request.Update_Stock_Rill

		stock_rill.Stock_rill = Request_detail.Stock_rill

		err := con.Table("detail_audit").Where("kode_detail_audit = ?", Request_detail.Kode_detail_audit).Select("stock_rill").Updates(&stock_rill)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_detail
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}

	}

	return res, nil
}

func Read_Audit_Stock(Request request.Read_Audit_Stock, Request_status request.Status_Audit_hari_ini_Request) (response.Response, error) {
	var res response.Response

	if Request_status.Status == 1 {
		var arr_data []response.Read_Audit_Stock_Response

		con := db.CreateConGorm().Table("stock")

		rows, err := con.Select("kode_stock", "nama_barang", "jumlah").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

		defer rows.Close()

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			//res.Data = data
			return res, err
		}

		for rows.Next() {
			var data response.Read_Audit_Stock_Response
			var detail_data []response.Detail_Audit_Stock_Response
			err = rows.Scan(&data.Kode_stock, &data.Nama_barang, &data.Total_jumlah_dalam_sistem)
			data.Tanggal = Request.Tanggal

			if err != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = data
				return res, err
			}

			con_blank := db.CreateConGorm()

			err2 := con_blank.Table("audit").Select("kode_audit").Where("kode_stock = ? && status = 0", data.Kode_stock).Scan(&data.Kode_audit)

			if err2.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = data
				return res, err2.Error
			}

			con_detail_barang := db.CreateConGorm().Table("detail_stock")

			if data.Kode_audit == "" {

				err2 = con_detail_barang.Select("detail_stock.kode_barang_keluar_masuk", "tanggal as tanggal_masuk", "jumlah_barang as stock_dalam_sistem", "skm.kode as kode_supplier").Joins("JOIN stock_keluar_masuk skm on skm.kode_stock_keluar_masuk = detail_stock.kode_stock_keluar_masuk").Joins("LEFT JOIN detail_audit da ON da.kode_barang_keluar_masuk = detail_stock.kode_barang_keluar_masuk").Where("kode_stock =? && detail_stock.jumlah_barang > 0", data.Kode_stock).Scan(&detail_data)

				if err2.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2.Error
				}

				for i := 0; i < len(detail_data); i++ {
					data.Total_jumlah_dalam_sistem = data.Total_jumlah_dalam_sistem + detail_data[i].Stock_dalam_sistem
				}

				data.Total_jumlah_selisih_stock = 0
				data.Total_jumlah_stock_rill = 0

			} else {

				err2 = con_detail_barang.Select("IFNULL(kode_detail_audit,'') AS kode_detail_audit", "detail_stock.kode_barang_keluar_masuk", "tanggal as tanggal_masuk", "jumlah_barang AS stock_dalam_sistem", "IFNULL(stock_rill, 0) AS stock_rill", "IFNULL(selisih_stock, 0) AS selisih_stock", "skm.kode as kode_supplier").Joins("JOIN stock_keluar_masuk skm on skm.kode_stock_keluar_masuk = detail_stock.kode_stock_keluar_masuk").Joins("LEFT JOIN detail_audit da ON da.kode_barang_keluar_masuk = detail_stock.kode_barang_keluar_masuk").Where("kode_stock = ? && detail_stock.jumlah_barang > 0", data.Kode_stock).Scan(&detail_data)

				if err2.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2.Error
				}

				con_detail_barang_upadate := db.CreateConGorm()

				var status request.Status_Audit_hari_ini_Request

				status.Status = 1

				var id []string

				err2 = con_detail_barang_upadate.Table("detail_stock").Select("kode_detail_audit").Joins("JOIN stock_keluar_masuk skm on skm.kode_stock_keluar_masuk = detail_stock.kode_stock_keluar_masuk").Joins("JOIN detail_audit da ON da.kode_barang_keluar_masuk = detail_stock.kode_barang_keluar_masuk").Where("kode_stock = ? && detail_stock.jumlah_barang = 0", data.Kode_stock).Scan(&id)

				fmt.Println(id)

				for i := 0; i < len(id); i++ {

					err := con_detail_barang_upadate.Table("detail_audit").Where("kode_detail_audit = ?", id[i]).Select("status").Updates(&status)

					if err.Error != nil {
						res.Status = http.StatusNotFound
						res.Message = "Status Not Found"
						res.Data = status
						return res, err.Error
					}

				}

				if err2.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2.Error
				}

				for i := 0; i < len(detail_data); i++ {

					detail_data[i].Selisih_stock = math.Abs(detail_data[i].Stock_dalam_sistem - detail_data[i].Stock_rill)

					data.Total_jumlah_dalam_sistem = data.Total_jumlah_dalam_sistem + detail_data[i].Stock_dalam_sistem
					data.Total_jumlah_stock_rill = data.Total_jumlah_stock_rill + detail_data[i].Stock_rill
				}

				data.Total_jumlah_selisih_stock = math.Abs(data.Total_jumlah_dalam_sistem - data.Total_jumlah_stock_rill)

			}

			data.Detail_audit_awal = detail_data

			if detail_data != nil {
				arr_data = append(arr_data, data)
			}

		}

		//fmt.Println(arr_data)

		if arr_data == nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = arr_data

		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = arr_data
		}

	} else if Request_status.Status == 0 {

		var data response.Read_Audit_Stock_Response
		var arr_data []response.Read_Audit_Stock_Response

		con := db.CreateConGorm().Table("audit")

		date, _ := time.Parse("02-01-2006", Request.Tanggal)
		Request.Tanggal = date.Format("2006-01-02")

		rows, err := con.Select("audit.kode_audit", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "audit.kode_stock", "nama_barang", "SUM(ds.stock_dalam_sistem)", "SUM(stock_rill)", "SUM(selisih_stock)").Joins("JOIN stock s ON s.kode_stock = audit.kode_stock").Joins("JOIN detail_audit ds ON ds.kode_audit = audit.kode_audit").Where("audit.kode_gudang = ? && tanggal = ? && audit.status = 1 && detail_audit.status = 0", Request.Kode_gudang, Request.Tanggal).Group("audit.kode_audit").Order("audit.co DESC").Rows()

		defer rows.Close()

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		for rows.Next() {
			var detail_data []response.Detail_Audit_Stock_Response
			err = rows.Scan(&data.Kode_audit, &data.Tanggal, &data.Kode_stock, &data.Nama_barang, &data.Total_jumlah_dalam_sistem, &data.Total_jumlah_stock_rill, &data.Total_jumlah_selisih_stock)

			if err != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = data
				return res, err
			}

			con_detail_audit := db.CreateConGorm().Table("detail_audit")

			err2 := con_detail_audit.Select("kode_detail_audit", "kode_barang_keluar_masuk", "DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock").Where("kode_audit = ?", data.Kode_audit).Scan(&detail_data)

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

	}

	return res, nil
}

// fungsi change status
func Update_Status_Audit(Request request.Update_Status_Audit_Request) (response.Response, error) {
	var res response.Response
	var err *gorm.DB

	con := db.CreateConGorm()

	kode_audit := tools.String_Separator_To_String(Request.Kode_audit)

	for i := 0; i < len(kode_audit); i++ {
		var status request.Update_Status_Audit_Request_V2

		status.Status = 1

		date, _ := time.Parse("02-01-2006", Request.Tanggal)
		status.Tanggal = date.Format("2006-01-02")

		err = con.Table("audit").Where("kode_audit = ?", kode_audit[i]).Select("status", "tanggal").Updates(&status)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		// pencatatan pada stock_keluar_masuk
		var Request_keluar request.Input_Stock_Keluar_Request
		con_stock_keluar := db.CreateConGorm().Table("stock_keluar_masuk")

		co := 0

		err = con_stock_keluar.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_keluar.Co = co + 1
		Request_keluar.Kode_stock_keluar_masuk = kode_audit[i]
		Request_keluar.Status = 2

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request_keluar.Tanggal = status.Tanggal
		Request_keluar.Kode_nota = ""

		con_user := db.CreateConGorm().Table("user")

		username := ""
		err = con_user.Select("username").Where("id_user = ?", Request.Kode_user).Scan(&username)

		Request_keluar.Nama_penanggung_jawab = username
		Request_keluar.Kode_gudang = Request.Kode_gudang
		Request_keluar.Kode = ""

		err = con_stock_keluar.Select("co", "kode_stock_keluar_masuk", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode", "kode_gudang", "status").Create(&Request_keluar)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		// adjust jumlah stock
		total_stock := 0.0
		err = con.Table("detail_audit").Select("SUM(stock_rill)").Where("kode_audit =?", kode_audit[i]).Scan(&total_stock)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		kode_stock := ""
		err = con.Table("audit").Select("kode_stock").Where("kode_audit =?", kode_audit[i]).Scan(&kode_stock)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		con_stock := db.CreateConGorm().Table("stock")

		err = con_stock.Where("kode_stock = ?", kode_stock).Update("jumlah", &total_stock)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		//adjust detail stock
		err = con.Raw("UPDATE detail_stock join detail_audit da on detail_stock.kode_detail_stock = da.kode_barang_keluar_masuk SET jumlah_barang = stock_rill WHERE da.kode_audit = ?", kode_audit[i])

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
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
