package retur_customer

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response"
	"POS-BACKEND/models/response_kasir"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_Retur_Customer(Request request_kasir.Input_Retur_Customer_Request, Request_Barang request_kasir.Input_Barang_Retur_Customer_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	con := db.CreateConGorm().Table("retur_customer")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_retur_customer = "RC-" + strconv.Itoa(Request.Co)

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

	err = con.Select("co", "kode_retur_customer", "kode_nota", "tanggal", "kode_bentuk_retur", "kode_store", "kode_kasir").Create(&Request)

	kode_barang_kasir := tools.String_Separator_To_String(Request_Barang.Kode_barang_kasir)
	jumlah := tools.String_Separator_To_float64(Request_Barang.Jumlah)
	keterangan := tools.String_Separator_To_String(Request_Barang.Keterangan)

	for i := 0; i < len(kode_barang_kasir); i++ {
		var barang_V2 request_kasir.Input_Barang_Retur_Customer_Request_V2

		con_barang := db.CreateConGorm().Table("barang_retur_customer")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_retur_customer = "BRC-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_retur_customer = Request.Kode_retur_customer
		barang_V2.Kode_barang_kasir = kode_barang_kasir[i]
		barang_V2.Jumlah = jumlah[i]
		barang_V2.Keterangan = keterangan[i]

		err = con_barang.Select("co", "kode_barang_retur_customer", "kode_retur_customer", "kode_barang_kasir", "jumlah", "keterangan").Create(&barang_V2)

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

func Read_Retur_Customer(Request request_kasir.Read_Retur_Customer_Request, Request_Filter request_kasir.Read_Filter_Retur_Customer_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	var arr_data []response_kasir.Read_Retur_Customer_Response
	var data response_kasir.Read_Retur_Customer_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT kode_retur_customer, kode_nota, DATE_FORMAT(tanggal, '%d-%m-%Y'), retur_customer.kode_bentuk_retur, nama_bentuk_retur , retur_customer.kode_store, nama_store, status FROM retur_customer join user_management um on um.kode_store = retur_customer.kode_store join bentuk_retur bp on bp.kode_bentuk_retur = retur_customer.kode_bentuk_retur WHERE retur_customer.kode_kasir = '" + Request.Kode_kasir + "'"

	if Request_Filter.Tanggal != "" {

		date, _ := time.Parse("02-01-2006", Request_Filter.Tanggal)
		Request_Filter.Tanggal = date.Format("2006-01-02")

		statement += " AND tanggal = '" + Request_Filter.Tanggal + "'"

	}

	if Request_Filter.Kode_store != "" {

		statement += " AND retur_customer.kode_store = '" + Request_Filter.Kode_store + "'"

	}

	statement += " ORDER BY retur_customer.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_retur_customer, &data.Kode_nota, &data.Tanggal, &data.Kode_bentuk_retur, &data.Nama_bentuk_retur, &data.Kode_store, &data.Nama_store, &data.Status)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_retur_customer")
		var detail_data []response_kasir.Detail_Retur_Customer_Response

		err = con_detail.Select("kode_barang_retur_customer", "barang_retur_customer.kode_barang_kasir", "nama_barang_kasir", "barang_retur_customer.jumlah", "keterangan").Joins("join barang_kasir bk on bk.kode_barang_kasir = barang_retur_customer.kode_barang_kasir").Where("kode_retur_customer = ?", data.Kode_retur_customer).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_Retur_Customer = detail_data

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

func Update_Retur_Customer(Request request_kasir.Update_Retur_Customer_Request, Request_kode request_kasir.Update_Kode_Retur_Customer_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	check := -1
	con_check := db.CreateConGorm().Table("retur_customer")

	err := con_check.Select("status").Joins("JOIN barang_retur_customer brc ON brc.kode_retur_customer = retur_customer.kode_retur_customer ").Where("kode_barang_retur_customer = ?", Request_kode.Kode_barang_retur_customer).Scan(&check)

	fmt.Println(check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 2 {

		con := db.CreateConGorm().Table("barang_retur_customer")

		err = con.Where("kode_barang_retur_customer = ?", Request_kode.Kode_barang_retur_customer).Select("kode_barang_kasir", "keterangan", "jumlah").Updates(&Request)

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

func Delete_Request_Barang_Kasir(Request request_kasir.Update_Kode_Retur_Customer_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	check := -1
	con_check := db.CreateConGorm().Table("retur_customer")

	err := con_check.Select("status").Joins("JOIN barang_retur_customer brc ON brc.kode_retur_customer = retur_customer.kode_retur_customer ").Where("kode_barang_retur_customer = ?", Request.Kode_barang_retur_customer).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 2 {

		con := db.CreateConGorm().Table("retur_customer")

		data := ""

		err = con.Select("retur_customer.kode_retur_customer").Joins("JOIN barang_retur_customer brc ON brc.kode_retur_customer = retur_customer.kode_retur_customer").Statement.Where("kode_barang_retur_customer = ?", Request.Kode_barang_retur_customer).Scan(&data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang := db.CreateConGorm().Table("barang_retur_customer")

		err = con_barang.Where("kode_barang_retur_customer = ?", Request.Kode_barang_retur_customer).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Delete Error"
			res.Data = Request
			return res, err.Error
		}

		kode_barang := ""

		con_check := db.CreateConGorm().Table("barang_retur_customer")

		err = con_check.Select("kode_barang_retur_customer").Where("kode_retur_customer=?", data).Limit(1).Scan(&kode_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Delete Error"
			res.Data = Request
			return res, err.Error
		}

		if kode_barang == "" {

			con_del_req := db.CreateConGorm().Table("retur_customer")

			err = con_del_req.Where("kode_retur_customer = ?", data).Delete("")

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

func Update_Status_Request_Barang_Kasir(Request request_kasir.Update_Status_Retur_Customer_Request, Request_kode request_kasir.Update_Status_Retur_Customer_Kode_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("retur_customer")
	status := -1

	err := con.Select("status").Where("kode_retur_customer = ?", Request_kode.Kode_retur_customer).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Status == 2 || Request.Status == 0 {

			con := db.CreateConGorm().Table("retur_customer")

			err := con.Where("kode_retur_customer = ?", Request_kode.Kode_retur_customer).Select("status").Updates(&Request)

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
			con := db.CreateConGorm()

			err := con.Table("retur_customer").Where("kode_retur_customer = ?", Request_kode.Kode_retur_customer).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var arr_barang []response_kasir.Detail_Retur_Customer_Response

			fmt.Println(Request_kode.Kode_retur_customer)

			br := ""

			err = con.Table("retur_customer").Select("kode_bentuk_retur").Where("kode_retur_customer = ?", Request_kode.Kode_retur_customer).Scan(&br)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			fmt.Println(br)

			if br == "BR-2" {

				err = con.Table("barang_retur_customer").Select("kode_barang_retur_customer", "barang_retur_customer.kode_barang_kasir", "nama_barang_kasir", "barang_retur_customer.jumlah", "keterangan").Joins("join barang_kasir bk on bk.kode_barang_kasir = barang_retur_customer.kode_barang_kasir").Where("kode_retur_customer = ?", Request_kode.Kode_retur_customer).Scan(&arr_barang)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				for i := 0; i < len(arr_barang); i++ {

					jumlah_lama := 0.0

					err = con.Table("barang_kasir").Select("jumlah").Where("kode_barang_kasir = ?", arr_barang[i].Kode_barang_kasir).Scan(&jumlah_lama)

					var jumlah request_kasir.Update_Jumlah_Barang_Kasir
					jumlah.Jumlah = arr_barang[i].Jumlah + jumlah_lama

					err = con.Table("barang_kasir").Where("kode_barang_kasir = ?", arr_barang[i].Kode_barang_kasir).Select("jumlah").Updates(&jumlah)

					if err.Error != nil {
						res.Status = http.StatusNotFound
						res.Message = "Status Not Found"
						res.Data = Request
						return res, err.Error
					}
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

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}
	return res, nil
}

func Dropdown_Kode_Nota_Retur_Customer(Request request_kasir.Read_Dropdown_Kode_Nota_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	var arr_data []response_kasir.Read_Dropdown_Kode_Nota_Request

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	con := db.CreateConGorm().Table("pembayaran")

	err := con.Select("kode_nota", "nama_store").Joins("join user_management um on um.kode_store = pembayaran.kode_store").Where("pembayaran.kode_kasir = ? AND tanggal = ?", Request.Kode_kasir, Request.Tanggal).Order("pembayaran.co ASC").Scan(&arr_data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
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

func Dropdown_Barang_Kasir_Retur(Request request_kasir.Read_Dropdown_Barang_Retur_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	var arr_data []response_kasir.Read_Detail_Pembayaran_Response

	con := db.CreateConGorm().Table("pembayaran")

	err := con.Select("kode_barang_pembayaran", "kode_barang_kasir", "nama_barang_kasir", "jumlah_barang", "nama_satuan", "harga").Joins("join barang_pembayaran bp on bp.kode_pembayaran = pembayaran.kode_pembayaran").Where("pembayaran.kode_nota", Request.Kode_nota).Order("pembayaran.co ASC").Scan(&arr_data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err
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
