package pre_order

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/services/stock_masuk"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Input_Pre_Order(Request request.Input_Pre_Order_Request, Request_Barang request.Input_Barang_Pre_Order_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("pre_order")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_pre_order = "PO-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	//0 = pending
	//1 = success
	//2 = ditolak

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")
	Request.Status = 0

	err = con.Select("co", "kode_pre_order", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode_supplier", "kode_gudang", "status").Create(&Request)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	Jumlah_barang := tools.String_Separator_To_float64(Request_Barang.Jumlah_barang)
	harga_pokok := tools.String_Separator_To_Int64(Request_Barang.Harga_pokok)
	tgl_kadaluarsa := tools.String_Separator_To_String(Request_Barang.Tanggal_kadalurasa)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Pre_Order_V2_Request

		con_barang := db.CreateConGorm().Table("barang_pre_order")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_pre_order = "BPO-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_pre_order = Request.Kode_pre_order
		barang_V2.Kode_stock = kode_stock[i]
		barang_V2.Jumlah_barang = Jumlah_barang[i]
		barang_V2.Harga = harga_pokok[i]
		barang_V2.Total_harga = int64(math.Round(float64(harga_pokok[i]) * Jumlah_barang[i]))

		date3, _ := time.Parse("02-01-2006", tgl_kadaluarsa[i])
		barang_V2.Tanggal_kadaluarsa = date3.Format("2006-01-02")

		fmt.Println(barang_V2.Tanggal_kadaluarsa)
		fmt.Println(barang_V2)

		err = con_barang.Select("co", "kode_barang_pre_order", "kode_stock", "kode_pre_order", "tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
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

func Read_Pre_Order(Request request.Read_Pre_Order_Request, Request_filter request.Read_Pre_Order_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Pre_Order_Response
	var data response.Read_Pre_Order_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT pre_order.kode_pre_order, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, kode_nota, nama_supplier, nama_penanggung_jawab, sum(jumlah_barang), sum(total_harga), status FROM pre_order JOIN supplier s ON s.kode_supplier = pre_order.kode_supplier JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order WHERE pre_order.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Kode_supplier != "" {
		statement += " AND kode_supplier = '" + Request_filter.Kode_supplier + "'"
	}

	if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" {

		statement += " AND (tanggal >= '" + Request_filter.Tanggal_1 + "' && tanggal <= '" + Request_filter.Tanggal_2 + "' )"

	} else if Request_filter.Tanggal_1 != "" {

		statement += " && tanggal = '" + Request_filter.Tanggal_1 + "'"

	}

	statement += " GROUP BY pre_order.kode_pre_order ORDER BY pre_order.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_pre_order, &data.Tanggal, &data.Kode_nota, &data.Penanggung_jawab, &data.Nama_supplier, &data.Jumlah_total, &data.Total_harga, &data.Status)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_pre_order")
		var detail_data []response.Read_Detail_Stock_Masuk_Response

		err = con_detail.Select("kode_barang_pre_order", "nama_barang", "DATE_FORMAT(tanggal_kadaluarsa, '%d-%m-%Y') AS tanggal_kadaluarsa", "jumlah_barang", "harga").Joins("join stock s on barang_pre_order.kode_stock = s.kode_stock").Where("kode_pre_order = ?", data.Kode_pre_order).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_stock_masuk = detail_data

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

func Update_Pre_Order(Request request.Update_Pre_order_Request, Request_kode request.Update_Pre_Order_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("pre_order")

	err := con_check.Select("status").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request_kode.Kode_barang_pre_order).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 {

		con := db.CreateConGorm().Table("barang_pre_order")

		date3, _ := time.Parse("02-01-2006", Request.Tanggal_kadaluarsa)
		Request.Tanggal_kadaluarsa = date3.Format("2006-01-02")

		Request.Total_harga = int64(math.Round(float64(Request.Harga) * Request.Jumlah_barang))

		err = con.Where("kode_barang_pre_order = ?", Request_kode.Kode_barang_pre_order).Select("tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga").Updates(&Request)

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
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Delete_Pre_Order(Request request.Update_Pre_Order_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("pre_order")

	err := con_check.Select("status").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 {

		con := db.CreateConGorm().Table("pre_order")

		data := ""

		err = con.Select("kode_pre_order").Where("kode_barang_pre_order=?", Request.Kode_barang_pre_order).Scan(&data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang := db.CreateConGorm().Table("barang_pre_order")

		err = con_barang.Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Delete("")

		kode_barang := ""

		err = con_barang.Select("kode_barang_pre_order").Where("kode_pre_order=?", data).Limit(1).Scan(&kode_barang)

		if kode_barang == "" {

			err = con.Where("kode_pre_order = ?", Request.Kode_barang_pre_order).Delete("")

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Update_Status_Pre_Order(Request request.Update_Status_Pre_Order_Request, Request_kode request.Kode_Pre_Order_Request) (response.Response, error) {
	var res response.Response
	var err2 error
	con := db.CreateConGorm().Table("pre_order")
	status := -1

	err := con.Select("status").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Status == 2 || Request.Status == 0 {

			con := db.CreateConGorm().Table("pre_order")

			err := con.Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Select("status").Updates(&Request)

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
		} else if Request.Status == 1 {
			con := db.CreateConGorm().Table("pre_order")

			err := con.Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var Request_Input_stock request.Input_Stock_Masuk_Request

			err = con.Select("tanggal", "kode_nota", "nama_penanggung_jawab", "kode_gudang").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&Request_Input_stock)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			err = con.Select("kode_supplier").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&Request_Input_stock.Kode)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			fmt.Println(Request_Input_stock.Kode)
			fmt.Println(Request_Input_stock)

			var Request_barang []request.Move_Barang_Pre_Order_Request

			con_barang := db.CreateConGorm().Table("barang_pre_order")
			err = con_barang.Select("kode_stock", "tanggal_kadaluarsa", "jumlah_barang", "harga").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&Request_barang)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request_barang
				return res, err.Error
			}

			fmt.Println(Request_barang)

			kode_stock := ""
			tanggal_kadaluarsa := ""
			jumlah_barang := ""
			harga := ""

			for i := 0; i < len(Request_barang); i++ {
				kode_stock = kode_stock + "|" + Request_barang[i].Kode_stock + "|"

				date, _ := time.Parse("2006-01-02", Request_barang[i].Tanggal_kadaluarsa)
				Request_barang[i].Tanggal_kadaluarsa = date.Format("02-01-2006")

				tanggal_kadaluarsa = tanggal_kadaluarsa + "|" + Request_barang[i].Tanggal_kadaluarsa + "|"

				jumlah_barang = jumlah_barang + "|" + fmt.Sprintf("%f", Request_barang[i].Jumlah_barang) + "|"

				harga = harga + "|" + strconv.FormatInt(Request_barang[i].Harga, 10) + "|"
			}

			var Request_Barang_V2 request.Input_Barang_Stock_Masuk_Request

			Request_Barang_V2.Kode_stock = kode_stock
			Request_Barang_V2.Tanggal_kadalurasa = tanggal_kadaluarsa
			Request_Barang_V2.Jumlah_barang = jumlah_barang
			Request_Barang_V2.Harga_pokok = harga

			fmt.Println(Request_Barang_V2)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			res, err2 = stock_masuk.Input_Stock_Masuk(Request_Input_stock, Request_Barang_V2)

			if err2 != nil {
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

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}
	return res, nil
}
