package request_barang_gudang

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response"
	"POS-BACKEND/services/gudang/stock_keluar"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Read_Request_Barang_Kasir_Stock(Request request.Read_Request_Barang_Kasir_Stock_Request, Request_filter request.Read_Filter_Request_Barang_Kasir_Stock_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Request_Barang_Kasir_Stock_Response
	var data response.Read_Request_Barang_Kasir_Stock_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT request_barang_kasir.kode_request_barang_kasir, DATE_FORMAT(tanggal_request, '%d-%m-%Y'), um.kode_store, nama_store, status, SUM(jumlah) AS jumlah FROM request_barang_kasir JOIN user_management um on um.kode_store = request_barang_kasir.kode_store JOIN gudang on gudang.kode_gudang = request_barang_kasir.kode_gudang JOIN toko tk on tk.kode_store = request_barang_kasir.kode_store JOIN barang_request_barang_kasir brbk on brbk.kode_request_barang_kasir = request_barang_kasir.kode_request_barang_kasir WHERE request_barang_kasir.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Tanggal_1 != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		Request_filter.Tanggal_1 = date.Format("2006-01-02")

		statement += " && tanggal_request = '" + Request_filter.Tanggal_1 + "'"

	}

	statement += "GROUP BY request_barang_kasir.kode_request_barang_kasir ORDER BY request_barang_kasir.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_request_barang_kasir, &data.Tanggal_request, &data.Kode_toko, &data.Nama_toko, &data.Status, &data.Jumlah)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_request_barang_kasir")
		var detail_data []response.Read_Barang_Request_Barang_Kasir_Stock_Response

		err = con_detail.Select("kode_barang_request", "kode_stock_gudang", "stk.nama_barang", "barang_request_barang_kasir.jumlah").Joins("join barang_kasir bk on barang_request_barang_kasir.kode_barang_kasir = bk.kode_barang_kasir").Joins("join stock stk on barang_request_barang_kasir.kode_stock_gudang = stk.kode_stock").Where("kode_request_barang_kasir = ?", data.Kode_request_barang_kasir).Scan(&detail_data).Error

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

func Update_Status_Request_Barang_Kasir(Request request_kasir.Update_Status_Request_Barang_Kasir, Request_kode request.Update_Status_Kode_Request_Barang_Request) (response.Response, error) {
	var res response.Response
	//var err2 error

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
		if Request.Status == 1 || Request.Status == 4 {
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

		} else if Request.Status == 2 {
			con := db.CreateConGorm().Table("request_barang_kasir")

			date := time.Now()
			tanggal_sekarang := date.Format("2006-01-02")

			err := con.Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var Request_Input_stock request.Input_Stock_Keluar_Request

			con_user := db.CreateConGorm().Table("user")

			penanggung_jawab := ""

			err = con_user.Select("username").Where("id_user = ?", Request_kode.Kode_user).Scan(&penanggung_jawab)

			Request_Input_stock.Kode_gudang = Request_kode.Kode_gudang
			Request_Input_stock.Tanggal = tanggal_sekarang
			Request_Input_stock.Nama_penanggung_jawab = penanggung_jawab
			Request_Input_stock.Kode_nota = Request_kode.Kode_nota

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var data_kasir response.Kode_Kasir_Kode_Store_Response

			err = con.Select("kode_kasir", "kode_store").Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Scan(&data_kasir)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			kode_toko := ""

			con_toko := db.CreateConGorm()

			err = con_toko.Table("toko").Select("kode_toko").Where("kode_store = ? AND kode_kasir = ?", data_kasir.Kode_store, data_kasir.Kode_kasir).Scan(&kode_toko)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			fmt.Println(kode_toko)

			if kode_toko == "" {
				var Request_toko request.Input_Toko_Request_Barang_Request

				err = con_toko.Table("user_management").Select("nama_store").Where("kode_store = ? AND kode_kasir = ?", data_kasir.Kode_store, data_kasir.Kode_kasir).Scan(&Request_toko.Nama_toko)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				co := 0

				err := con_toko.Table("toko").Select("co").Order("co DESC").Limit(1).Scan(&co)

				Request_toko.Co = co + 1
				Request_toko.Kode_toko = "TK-" + strconv.Itoa(Request_toko.Co)
				Request_toko.Kode_gudang = Request_kode.Kode_gudang
				Request_toko.Alamat = ""
				Request_toko.Nomor_telpon = ""
				Request_toko.Kode_kasir = data_kasir.Kode_kasir
				Request_toko.Kode_store = data_kasir.Kode_store

				fmt.Println(co)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = co
					return res, err.Error
				}

				err = con_toko.Table("toko").Select("co", "kode_toko", "nama_toko", "alamat", "nomor_telpon", "kode_gudang", "kode_kasir", "kode_store").Create(&Request_toko)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				Request_Input_stock.Kode = Request_toko.Kode_toko

			} else {
				Request_Input_stock.Kode = kode_toko
			}

			fmt.Println(Request_Input_stock.Kode)

			con_barang := db.CreateConGorm().Table("barang_request_barang_kasir")
			rows, err2 := con_barang.Select("kode_stock_gudang", "jumlah").Where("kode_request_barang_kasir = ?", Request_kode.Kode_request_barang_kasir).Rows()

			defer rows.Close()

			if err2 != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err2
			}

			kode_stock := ""
			jumlah_barang := ""
			harga := ""

			for rows.Next() {
				var data request.Move_Refund_To_Stock_Keluar_Request
				err2 = rows.Scan(&data.Kode_stock, &data.Jumlah_barang)

				if err2 != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2
				}
				kode_stock = kode_stock + "|" + data.Kode_stock + "|"

				jumlah_barang = jumlah_barang + "|" + fmt.Sprintf("%f", data.Jumlah_barang) + "|"

				harga = harga + "|" + "0" + "|"
			}

			var Request_Barang_V2 request.Input_Barang_Stock_Keluar_Request

			Request_Barang_V2.Kode_stock = kode_stock
			Request_Barang_V2.Jumlah_barang = jumlah_barang
			Request_Barang_V2.Harga_jual = harga

			fmt.Println(Request_Barang_V2)
			fmt.Println(Request_Input_stock)

			res, err2 = stock_keluar.Input_Stock_Keluar(Request_Input_stock, Request_Barang_V2)

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
