package refund_supplier

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"

	"POS-BACKEND/services/stock_keluar"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_Refund_Supplier(Request request.Input_Refund_Request, Request_Barang request.Input_Barang_Refund_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("refund")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_refund = "RF-" + strconv.Itoa(Request.Co)

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

	Request.Tanggal_pengembalian = "0001-01-01"

	err = con.Select("co", "kode_refund", "tanggal", "tanggal_pengembalian", "kode_supplier", "kode_gudang").Create(&Request)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	jumlah := tools.String_Separator_To_float64(Request_Barang.Jumlah)
	kode_nota := tools.String_Separator_To_String(Request_Barang.Kode_nota)
	keterangan := tools.String_Separator_To_String(Request_Barang.Keterangan)
	tgl_stock_masuk := tools.String_Separator_To_String(Request_Barang.Tanggal_stock_masuk)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Refund_V2_Request

		con_barang := db.CreateConGorm().Table("barang_refund")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_refund = "BRF-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_refund = Request.Kode_refund
		barang_V2.Kode_stock = kode_stock[i]
		barang_V2.Jumlah = jumlah[i]
		barang_V2.Kode_nota = kode_nota[i]
		barang_V2.Keterangan = keterangan[i]

		date3, _ := time.Parse("02-01-2006", tgl_stock_masuk[i])
		barang_V2.Tanggal_stock_masuk = date3.Format("2006-01-02")

		err = con_barang.Select("co", "kode_barang_refund", "kode_nota", "kode_stock", "tanggal_stock_masuk", "jumlah", "keterangan", "kode_refund").Create(&barang_V2)

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

func Read_Refund(Request request.Read_Refund_Request, Request_filter request.Read_Refund_Filter_Request) (response.Response, error) {
	var res response.Response

	var arr_data []response.Read_Refund_Response
	var data response.Read_Refund_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT refund.kode_refund, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, DATE_FORMAT(tanggal_pengembalian, '%d-%m-%Y') AS tanggal_pengembalian, nama_supplier, sum(jumlah), status FROM refund JOIN supplier s ON s.kode_supplier = refund.kode_supplier JOIN barang_refund br ON br.kode_refund = refund.kode_refund WHERE refund.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Kode_supplier != "" {
		statement += " AND refund.kode_supplier = '" + Request_filter.Kode_supplier + "'"
	}

	if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" {

		statement += " AND (tanggal >= '" + Request_filter.Tanggal_1 + "'" + " && tanggal <= '" + Request_filter.Tanggal_2 + "' )"

	} else if Request_filter.Tanggal_1 != "" {

		statement += " && tanggal = '" + Request_filter.Tanggal_1 + "'"

	}

	statement += " GROUP BY refund.kode_refund ORDER BY refund.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_refund, &data.Tanggal, &data.Tanggal_pengambalian, &data.Nama_supplier, &data.Jumlah_total, &data.Status)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_refund")
		var detail_data []response.Read_Barang_Refund_Response

		err = con_detail.Select("kode_nota", "nama_barang", "DATE_FORMAT(tanggal_stock_masuk, '%d-%m-%Y') AS tanggal_stock_masuk", "barang_refund.jumlah", "keterangan").Joins("join stock s on barang_refund.kode_stock = s.kode_stock").Where("kode_refund = ?", data.Kode_refund).Scan(&detail_data).Error

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

func Update_Barang_Refund(Request request.Update_Refund_Request, Request_Barang request.Update_Barang_Refund_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("refund")

	err := con_check.Select("status").Joins("JOIN barang_refund br ON br.kode_refund = refund.kode_refund ").Where("kode_barang_refuns = ?", Request.Kode_barang_refund).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == 0 {

		con := db.CreateConGorm().Table("barang_refund")

		date3, _ := time.Parse("02-01-2006", Request_Barang.Tanggal_stock_masuk)
		Request_Barang.Tanggal_stock_masuk = date3.Format("2006-01-02")

		err = con.Where("kode_barang_refund = ?", Request.Kode_barang_refund).Select("kode_nota", "kode_stock", "tanggal_stock_masuk", "jumlah", "keterangan").Updates(&Request_Barang)

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

func Delete_Barang_Refund(Request request.Update_Refund_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("refund")

	err := con_check.Select("status").Joins("JOIN barang_refund br ON br.kode_refund = refund.kode_refund ").Where("br.kode_barang_refund = ?", Request.Kode_barang_refund).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	fmt.Println(check)
	fmt.Println(Request)

	if check == 0 {

		con := db.CreateConGorm().Table("refund")

		data := ""

		err = con.Select("refund.kode_refund").Joins("JOIN barang_refund br ON br.kode_refund = refund.kode_refund ").Where("kode_barang_refund=?", Request.Kode_barang_refund).Scan(&data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang := db.CreateConGorm().Table("barang_refund")

		err = con_barang.Where("kode_barang_refund = ?", Request.Kode_barang_refund).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Delete Error"
			res.Data = Request
			return res, err.Error
		}

		kode_barang := ""

		con_barang2 := db.CreateConGorm().Table("barang_refund")

		err = con_barang2.Select("kode_barang_refund").Where("kode_refund=?", data).Limit(1).Scan(&kode_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "SELECT error"
			res.Data = Request
			return res, err.Error
		}

		if kode_barang == "" {

			con_refund := db.CreateConGorm().Table("refund")

			err = con_refund.Where("kode_refund = ?", data).Delete("")

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

func Update_Status_Refund(Request request.Update_Status_Refund_Request, Request_kode request.Update_Status_Kode_Refund_Request) (response.Response, error) {
	var res response.Response
	//var err2 error
	con := db.CreateConGorm().Table("refund")
	status := -1

	err := con.Select("status").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Status == 2 || Request.Status == 0 {

			con := db.CreateConGorm().Table("refund")

			err := con.Where("kode_refund = ?", Request_kode.Kode_refund).Select("status").Updates(&Request)

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
			var data_update request.Data_Update_Refund
			con := db.CreateConGorm().Table("refund")

			date := time.Now()
			tanggal_sekarang := date.Format("2006-01-02")

			fmt.Println(tanggal_sekarang)
			data_update.Status = Request.Status
			data_update.Tanggal_pengembalian = tanggal_sekarang

			err := con.Where("kode_refund = ?", Request_kode.Kode_refund).Select("tanggal_pengembalian", "status").Updates(&data_update)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var Request_Input_stock request.Input_Stock_Keluar_Request

			var Temp request.Tanggal_dan_Kode_Gudang

			con_user := db.CreateConGorm().Table("user")

			penanggung_jawab := ""

			err = con_user.Select("username").Where("id_user = ?", Request_kode.Kode_user).Scan(&penanggung_jawab)

			err = con.Select("tanggal_pengembalian", "kode_gudang").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&Temp)

			Request_Input_stock.Kode_gudang = Temp.Kode_gudang
			Request_Input_stock.Tanggal = Temp.Tanggal_pengembalian
			Request_Input_stock.Nama_penanggung_jawab = penanggung_jawab

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			err = con.Select("kode_supplier").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&Request_Input_stock.Kode)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			fmt.Println(Request_Input_stock.Kode)

			con_barang := db.CreateConGorm().Table("barang_refund")
			rows, err2 := con_barang.Select("kode_nota", "kode_stock", "jumlah").Where("kode_refund = ?", Request_kode.Kode_refund).Rows()

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
				err2 = rows.Scan(&Request_Input_stock.Kode_nota, &data.Kode_stock, &data.Jumlah_barang)
				fmt.Println(Request_Input_stock)

				if err2 != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2
				}
				kode_stock = kode_stock + "|" + data.Kode_stock + "|"

				jumlah_barang = jumlah_barang + "|" + fmt.Sprintf("%f", data.Jumlah_barang) + "|"

				harga = harga + "|" + "0" + "|"

				var Request_Barang_V2 request.Input_Barang_Stock_Keluar_Request

				Request_Barang_V2.Kode_stock = kode_stock
				Request_Barang_V2.Jumlah_barang = jumlah_barang
				Request_Barang_V2.Harga_jual = harga

				fmt.Println(Request_Barang_V2)

				res, err2 = stock_keluar.Input_Stock_Keluar(Request_Input_stock, Request_Barang_V2)
			}

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
