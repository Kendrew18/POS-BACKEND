package request_barang_kasir

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response"
	"POS-BACKEND/models/response_kasir"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Input_Request_Barang_Kasir(Request request_kasir.Input_Request_Barang_Kasir_Request, Request_Barang request_kasir.Input_Barang_Request) (response_kasir.Response, error) {

	var res response_kasir.Response

	con := db.CreateConGorm().Table("request_barang_kasir")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_request_barang_kasir = "RBK-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	//0 = pending (Kasir)
	//1 = proses -> tidak bisa di edit (Gudang)
	//2 = dikirim -> pengurangan stock gudang (Gudang)
	//3 = sukses -> sudah di terima (Kasir)
	//4 = ditolak -> (dua-duanya)

	date, _ := time.Parse("02-01-2006", Request.Tanggal_request)
	Request.Tanggal_request = date.Format("2006-01-02")
	Request.Status = 0

	err = con.Select("co", "kode_request_barang_kasir", "tanggal_request", "kode_gudang_kasir", "kode_store", "kode_kasir", "status").Create(&Request)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock_gudang)
	Jumlah := tools.String_Separator_To_float64(Request_Barang.Jumlah)
	kode_barang_kasir := tools.String_Separator_To_String(Request_Barang.Kode_barang_kasir)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request_kasir.Input_Barang_Request_V2

		con_barang := db.CreateConGorm().Table("barang_request_barang_kasir")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_request = "BRB-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_request_barang_kasir = Request.Kode_request_barang_kasir
		barang_V2.Kode_stock_gudang = kode_stock[i]
		barang_V2.Jumlah = Jumlah[i]
		barang_V2.Kode_barang_kasir = kode_barang_kasir[i]

		fmt.Println(barang_V2)

		err = con_barang.Select("co", "kode_barang_request", "kode_request_barang_kasir", "kode_stock_gudang", "jumlah", "kode_barang_kasir").Create(&barang_V2)

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

func Read_Request_Barang_Kasir(Request request_kasir.Read_Request_Barang_Kasir_Request, Request_filter request_kasir.Read_Filter_Request_Barang_Kasir) (response_kasir.Response, error) {

	var res response_kasir.Response
	var arr_data []response_kasir.Read_Request_Barang_Kasir_Response
	var data response_kasir.Read_Request_Barang_Kasir_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT request_barang_kasir.kode_request_barang_kasir, DATE_FORMAT(tanggal_request, '%d-%m-%Y'), request_barang_kasir.kode_gudang_kasir, GK.kode_gudang, nama_gudang, um.kode_store, nama_store, status FROM request_barang_kasir JOIN user_management um on um.kode_store = request_barang_kasir.kode_store JOIN gudang_kasir GK on GK.kode_gudang_kasir = request_barang_kasir.kode_gudang_kasir JOIN gudang on gudang.kode_gudang = GK.kode_gudang WHERE request_barang_kasir.kode_kasir = '" + Request.Kode_kasir + "'"

	if Request_filter.Kode_store != "" {
		statement += " && request_barang_kasir.kode_store = '" + Request_filter.Kode_store + "'"
	}

	if Request_filter.Tanggal_1 != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		Request_filter.Tanggal_1 = date.Format("2006-01-02")

		statement += " && tanggal_request = '" + Request_filter.Tanggal_1 + "'"

	}

	statement += " ORDER BY request_barang_kasir.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_request_barang_kasir, &data.Tanggal_request, &data.Kode_gudang_kasir, &data.Kode_gudang, &data.Nama_gudang, &data.Kode_store, &data.Nama_store, &data.Status)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_request_barang_kasir")
		var detail_data []response_kasir.Read_Barang_Request_Barang_Kasir_Response

		err = con_detail.Select("kode_barang_request", "barang_request_barang_kasir.kode_barang_kasir", "nama_barang_kasir", "kode_stock_gudang", "stk.nama_barang", "barang_request_barang_kasir.jumlah").Joins("join barang_kasir bk on barang_request_barang_kasir.kode_barang_kasir = bk.kode_barang_kasir").Joins("join stock stk on barang_request_barang_kasir.kode_stock_gudang = stk.kode_stock").Where("kode_request_barang_kasir = ?", data.Kode_request_barang_kasir).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_barang = detail_data

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

func Update_Request_Barang_Kasir(Request request_kasir.Update_Request_Barang_Kasir_Request, Request_kode request_kasir.Update_Request_Barang_Kasir_Kode) (response_kasir.Response, error) {
	var res response_kasir.Response

	check := -1
	con_check := db.CreateConGorm().Table("request_barang_kasir")

	err := con_check.Select("status").Joins("JOIN barang_request_barang_kasir brbk ON brbk.kode_request_barang_kasir = request_barang_kasir.kode_request_barang_kasir ").Where("kode_barang_request = ?", Request_kode.Kode_barang_request).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 4 {

		con := db.CreateConGorm().Table("barang_request_barang_kasir")

		err = con.Where("kode_barang_request = ?", Request_kode.Kode_barang_request).Select("kode_stock_gudang", "kode_barang_kasir", "jumlah").Updates(&Request)

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

func Delete_Request_Barang_Kasir(Request request_kasir.Update_Request_Barang_Kasir_Kode) (response_kasir.Response, error) {
	var res response_kasir.Response

	check := -1
	con_check := db.CreateConGorm().Table("request_barang_kasir")

	err := con_check.Select("status").Joins("JOIN barang_request_barang_kasir brbk ON brbk.kode_request_barang_kasir = request_barang_kasir.kode_request_barang_kasir ").Where("kode_barang_request = ?", Request.Kode_barang_request).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 || check == 4 {

		con := db.CreateConGorm().Table("request_barang_kasir")

		data := ""

		err = con.Select("request_barang_kasir.kode_request_barang_kasir").Joins("JOIN barang_request_barang_kasir brbk ON brbk.kode_request_barang_kasir = request_barang_kasir.kode_request_barang_kasir").Statement.Where("kode_barang_request = ?", Request.Kode_barang_request).Scan(&data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang := db.CreateConGorm().Table("barang_request_barang_kasir")

		err = con_barang.Where("kode_barang_request = ?", Request.Kode_barang_request).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Delete Error"
			res.Data = Request
			return res, err.Error
		}

		kode_barang := ""

		con_check := db.CreateConGorm().Table("barang_request_barang_kasir")

		err = con_check.Select("kode_barang_request").Where("kode_request_barang_kasir=?", data).Limit(1).Scan(&kode_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Delete Error"
			res.Data = Request
			return res, err.Error
		}

		if kode_barang == "" {

			con_del_req := db.CreateConGorm().Table("request_barang_kasir")

			err = con_del_req.Where("kode_request_barang_kasir = ?", data).Delete("")

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

func Update_Status_Request_Barang_Kasir(Request request_kasir.Update_Status_Request_Barang_Kasir, Request_kode request_kasir.Kode_Request_Barang_Kasir_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("request_barang_kasir")
	status := -1

	err := con.Select("status").Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 3 {
		if Request.Status == 4 || Request.Status == 0 {

			con := db.CreateConGorm().Table("request_barang_kasir")

			err := con.Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Select("status").Updates(&Request)

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

		} else if Request.Status == 3 {
			con := db.CreateConGorm()

			err := con.Table("request_barang_kasir").Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var arr_barang []response_kasir.Read_Barang_Request_Kasir_Response

			fmt.Println(Request_kode.Kode_request_barang_kasir)

			err = con.Table("barang_request_barang_kasir").Select("barang_request_barang_kasir.kode_barang_kasir", "barang_request_barang_kasir.jumlah", "jumlah_pengali").Joins("JOIN barang_kasir bk ON bk.kode_barang_kasir = barang_request_barang_kasir.kode_barang_kasir").Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Scan(&arr_barang)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			for i := 0; i < len(arr_barang); i++ {
				var jumlah request_kasir.Update_Jumlah_Barang_Kasir
				jumlah.Jumlah = math.Round(arr_barang[i].Jumlah*arr_barang[i].Jumlah_pengali*100) / 100

				err = con.Table("barang_kasir").Where("kode_barang_kasir = ?", arr_barang[i].Kode_barang_kasir).Select("jumlah").Updates(&jumlah)

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

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}
	return res, nil
}

func Dropdown_status(Request request_kasir.Dropdown_Status_Kasir_Request) (response_kasir.Response, error) {
	var res response_kasir.Response
	var data response_kasir.Status_Request_Barang_Kasir_Response
	var arr_data []response_kasir.Status_Request_Barang_Kasir_Response

	if strings.HasPrefix(Request.Kode, "KG") {

		con := db.CreateConGorm()
		status := -1

		err := con.Table("request_barang_kasir").Select("status").Where("kode_request_barang_kasir = ?", Request.Kode_request_barang_kasir).Scan(&status)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		if status == 0 {
			data.Status = 1
			data.Nama_status = "Diproses"

			arr_data = append(arr_data, data)

			data.Status = 4
			data.Nama_status = "Ditolak / Cancel"

			arr_data = append(arr_data, data)
		} else if status == 1 {
			data.Status = 2
			data.Nama_status = "Dikirim"

			arr_data = append(arr_data, data)

			data.Status = 4
			data.Nama_status = "Ditolak / Cancel"

			arr_data = append(arr_data, data)
		} else {
			arr_data = append(arr_data, data)

			res.Status = http.StatusNotFound
			res.Message = "Tidak dapat mengubah status"
			res.Data = Request

			return res, nil
		}

	} else if strings.HasPrefix(Request.Kode, "KS") {

		con := db.CreateConGorm()
		status := -1

		err := con.Table("request_barang_kasir").Select("status").Where("kode_request_barang_kasir = ?", Request.Kode_request_barang_kasir).Scan(&status)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		if status == 0 {
			data.Status = 4
			data.Nama_status = "Ditolak / Cancel"

			arr_data = append(arr_data, data)
		} else if status == 2 {
			data.Status = 3
			data.Nama_status = "Sukses"

			arr_data = append(arr_data, data)
		} else if status == 4 {
			data.Status = 0
			data.Nama_status = "Pending"

			arr_data = append(arr_data, data)
		} else {
			arr_data = append(arr_data, data)

			res.Status = http.StatusNotFound
			res.Message = "Tidak dapat mengubah status"
			res.Data = Request

			return res, nil
		}

	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = arr_data

	return res, nil
}
