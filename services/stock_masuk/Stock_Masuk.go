package stock_masuk

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Input_Stock_Masuk(Request_stock_masuk request.Input_Stock_Masuk_Request, Request_Barang request.Input_Barang_Stock_Masuk_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("stock_keluar_masuk")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request_stock_masuk.Co = co + 1
	Request_stock_masuk.Kode_stock_keluar_masuk = "SM-" + strconv.Itoa(Request_stock_masuk.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request_stock_masuk
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request_stock_masuk.Tanggal)
	Request_stock_masuk.Tanggal = date.Format("2006-01-02")
	Request_stock_masuk.Status = 0

	err = con.Select("co", "kode_stock_keluar_masuk", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode", "kode_gudang", "status").Create(&Request_stock_masuk)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	Jumlah_barang := tools.String_Separator_To_float64(Request_Barang.Jumlah_barang)
	harga_pokok := tools.String_Separator_To_Int64(Request_Barang.Harga_pokok)
	tgl_kadaluarsa := tools.String_Separator_To_String(Request_Barang.Tanggal_kadalurasa)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Stock_Masuk_V2_Request

		con_barang := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar_masuk = "BKM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar_masuk = Request_stock_masuk.Kode_stock_keluar_masuk
		barang_V2.Kode_stock = kode_stock[i]
		barang_V2.Jumlah_barang = Jumlah_barang[i]
		barang_V2.Harga = harga_pokok[i]
		barang_V2.Total_harga = int64(math.Round(float64(harga_pokok[i]) * Jumlah_barang[i]))

		date3, _ := time.Parse("02-01-2006", tgl_kadaluarsa[i])
		barang_V2.Tanggal_kadaluarsa = date3.Format("2006-01-02")

		fmt.Println(barang_V2.Tanggal_kadaluarsa)
		fmt.Println(barang_V2)

		err = con_barang.Select("co", "kode_barang_keluar_masuk", "kode_stock", "kode_stock_keluar_masuk", "tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		con_detail_stock := db.CreateConGorm().Table("detail_stock")

		err = con_detail_stock.Select("co", "kode_barang_keluar_masuk", "kode_stock", "kode_stock_keluar_masuk", "tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga", "status").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		jumlah_lama := (0.0)

		err = db.CreateConGorm().Table("stock").Select("jumlah").Where("kode_stock =?", kode_stock[i]).Scan(&jumlah_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		jumlah_baru := jumlah_lama + barang_V2.Jumlah_barang

		err = db.CreateConGorm().Table("stock").Where("kode_stock = ?", barang_V2.Kode_stock).Update("jumlah", jumlah_baru)
	}

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request_stock_masuk
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

func Read_Stock_Masuk(Request request.Read_Stock_Masuk_Request, Request_filter request.Read_Stock_Masuk_Filter_Request) (response.Response, error) {

	var res response.Response

	var arr_data []response.Read_Stock_Masuk_Response
	var data response.Read_Stock_Masuk_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm().Table("stock_keluar_masuk")

	if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" && Request_filter.Kode_supplier != "" {
		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_2)
		date_sql2 := date2.Format("2006-01-02")

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk ").Where("stock_keluar_masuk.kode_gudang = ? && (tanggal >= ? && tanggal <= ?) && stock_keluar_masuk.kode = ? && status=0", Request.Kode_gudang, date_sql, date_sql2, Request_filter.Kode_supplier).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 == "" && Request_filter.Kode_supplier != "" {
		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk ").Where("stock_keluar_masuk.kode_gudang = ? && tanggal = ? && stock_keluar_masuk.kode = ? && status=0", Request.Kode_gudang, date_sql, Request_filter.Kode_supplier).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" && Request_filter.Kode_supplier == "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_2)
		date_sql2 := date2.Format("2006-01-02")

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("stock_keluar_masuk.kode_gudang = ? && (tanggal >= ? && tanggal <= ?) && status=0", Request.Kode_gudang, date_sql, date_sql2).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 == "" && Request_filter.Kode_supplier == "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("stock_keluar_masuk.kode_gudang = ? && tanggal = ? && status=0", Request.Kode_gudang, date_sql).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	} else if Request_filter.Kode_supplier != "" {

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("stock_keluar_masuk.kode_gudang = ? && kode = ? && status=0", Request.Kode_gudang, Request_filter.Kode_supplier).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	} else {

		rows, err = con.Select("stock_keluar_masuk.kode_stock_keluar_masuk", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN supplier s ON s.kode_supplier = stock_keluar_masuk.kode").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk ").Where("stock_keluar_masuk.kode_gudang = ? && status=0", Request.Kode_gudang).Group("stock_keluar_masuk.kode_stock_keluar_masuk").Order("stock_keluar_masuk.co ASC").Rows()

	}

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&data.Kode_stock_keluar_masuk, &data.Tanggal, &data.Kode_nota, &data.Penanggung_jawab, &data.Nama_supplier, &data.Jumlah_total, &data.Total_harga)
		con_detail := db.CreateConGorm().Table("barang_stock_keluar_masuk")
		var detail_data []response.Read_Detail_Stock_Masuk_Response

		err := con_detail.Select("kode_barang_keluar_masuk", "nama_barang", "DATE_FORMAT(tanggal_kadaluarsa, '%d-%m-%Y') AS tanggal_kadaluarsa", "jumlah_barang", "harga", "status").Joins("join stock s on barang_stock_keluar_masuk.kode_stock = s.kode_stock").Where("kode_stock_keluar_masuk = ?", data.Kode_stock_keluar_masuk).Scan(&detail_data).Error

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

func Update_Barang_Stock_Masuk(Request request.Update_Stock_Masuk_Request, Request_kode request.Update_Stock_Masuk_Kode_Request) (response.Response, error) {

	var res response.Response
	var data response.Stock_Masuk_Lama_Response
	check := ""
	con_check := db.CreateConGorm().Table("pengurangan_stock")

	err := con_check.Select("kode_pengurangan").Where("kode_barang_keluar_masuk = ?", Request_kode.Kode_barang_keluar_masuk).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == "" {
		con := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		err = con.Select("kode_stock_keluar_masuk", "kode_stock", "jumlah_barang").Where("kode_barang_keluar_masuk=?", Request_kode.Kode_barang_keluar_masuk).Scan(&data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		jumlah_lama := 0.0

		err = db.CreateConGorm().Table("stock").Select("jumlah").Where("kode_stock =?", data.Kode_stock).Scan(&jumlah_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		jumlah_baru := jumlah_lama - data.Jumlah_barang + Request.Jumlah_barang

		err = db.CreateConGorm().Table("stock").Where("kode_stock = ?", data.Kode_stock).Update("jumlah", jumlah_baru)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		date3, _ := time.Parse("02-01-2006", Request.Tanggal_kadaluarsa)
		Request.Tanggal_kadaluarsa = date3.Format("2006-01-02")

		Request.Total_harga = int64(math.Round(float64(Request.Harga) * Request.Jumlah_barang))

		err = con.Where("kode_barang_keluar_masuk = ?", Request_kode.Kode_barang_keluar_masuk).Select("tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga").Updates(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_detail_barang := db.CreateConGorm().Table("detail_stock")

		err = con_detail_barang.Where("kode_barang_keluar_masuk = ?", Request_kode.Kode_barang_keluar_masuk).Select("tanggal_kadaluarsa", "jumlah_barang", "harga", "total_harga").Updates(&Request)

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

func Delete_Barang_Stock_Masuk(Request request.Update_Stock_Masuk_Kode_Request) (response.Response, error) {

	var res response.Response
	var data response.Stock_Masuk_Lama_Response
	check := ""
	con_check := db.CreateConGorm().Table("pengurangan_stock")

	err := con_check.Select("kode_pengurangan").Where("kode_barang_keluar_masuk = ?", Request.Kode_barang_keluar_masuk).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == "" {
		con := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		err = con.Select("kode_stock_keluar_masuk", "kode_stock", "jumlah_barang").Where("kode_barang_keluar_masuk=?", Request.Kode_barang_keluar_masuk).Scan(&data)

		fmt.Println(data.Kode_stock)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		jumlah_lama := 0.0

		err = db.CreateConGorm().Table("stock").Select("jumlah").Where("kode_stock =?", data.Kode_stock).Scan(&jumlah_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		jumlah_baru := jumlah_lama - data.Jumlah_barang

		err = db.CreateConGorm().Table("stock").Where("kode_stock = ?", data.Kode_stock).Update("jumlah", jumlah_baru)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_detail_barang := db.CreateConGorm().Table("detail_stock")

		err = con_detail_barang.Where("kode_barang_keluar_masuk = ?", Request.Kode_barang_keluar_masuk).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang_keluar_masuk := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		err = con_barang_keluar_masuk.Where("kode_barang_keluar_masuk = ?", Request.Kode_barang_keluar_masuk).Delete("")

		kode_barang := ""

		err = con_barang_keluar_masuk.Select("kode_barang_keluar_masuk").Where("kode_stock_keluar_masuk=?", data.Kode_stock_keluar_masuk).Limit(1).Scan(&kode_barang)

		if kode_barang == "" {
			con_stock_keluar_masuk := db.CreateConGorm().Table("stock_keluar_masuk")

			err = con_stock_keluar_masuk.Where("kode_stock_keluar_masuk = ?", Request.Kode_barang_keluar_masuk).Delete("")

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
