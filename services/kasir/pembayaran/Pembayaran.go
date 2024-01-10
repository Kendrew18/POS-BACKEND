package pembayaran

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/models/response_kasir"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_Pembayaran(Request request_kasir.Input_Pembayaran_Request, Request_barang request_kasir.Barang_Input_Pembayaran_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	con := db.CreateConGorm()

	co := 0

	err := con.Table("pembayaran").Select("co").Order("co DESC").Limit(1).Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	Request.Co = co + 1
	Request.Kode_pembayaran = "PMB-" + strconv.Itoa(Request.Co)

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")
	Request.Kode_nota = date.Format("20060102") + "-"

	co_pembayaran := 0

	err = con.Table("pembayaran").Select("COUNT(co)").Where("tanggal = ?", Request.Tanggal).Order("co DESC").Scan(&co_pembayaran)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	co_pembayaran = co_pembayaran + 1

	Request.Kode_nota = Request.Kode_nota + strconv.Itoa(co_pembayaran)
	Request.Jumlah_total = 0.0
	Request.Harga_total = 0

	err = con.Table("pembayaran").Select("co", "kode_pembayaran", "kode_nota", "tanggal", "kode_jenis_pembayaran", "kode_store", "jumlah_total", "total_harga", "kode_kasir").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	tot_jumlah := 0.0
	tot_harga := int64(0)

	kode_stock := tools.String_Separator_To_String(Request_barang.Kode_barang_kasir)
	Jumlah := tools.String_Separator_To_float64(Request_barang.Jumlah_barang)
	nama_satuan := tools.String_Separator_To_String(Request_barang.Nama_satuan)
	harga := tools.String_Separator_To_Int64(Request_barang.Harga)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request_kasir.Barang_Input_Pembayaran_Request_V2

		con_barang := db.CreateConGorm().Table("barang_pembayaran")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_pembayaran = "BPM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_pembayaran = Request.Kode_pembayaran
		barang_V2.Kode_barang_kasir = kode_stock[i]

		err = con.Table("barang_kasir").Select("nama_barang_kasir").Where("kode_barang_kasir = ?", kode_stock[i]).Scan(&barang_V2.Nama_barang_kasir)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Jumlah_barang = Jumlah[i]
		barang_V2.Nama_satuan = nama_satuan[i]
		barang_V2.Harga = harga[i]

		fmt.Println(barang_V2)

		err = con_barang.Select("co", "kode_barang_pembayaran", "kode_pembayaran", "kode_barang_kasir", "nama_barang_kasir", "jumlah_barang", "nama_satuan", "harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		tot_jumlah = tot_jumlah + Jumlah[i]
		tot_harga = tot_harga + harga[i]

		stock_lama := 0.0

		err = con.Table("barang_kasir").Select("jumlah").Where("kode_barang_kasir = ?", kode_stock[i]).Scan(&stock_lama)

		//fmt.Println(stock_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		stock_baru := stock_lama - Jumlah[i]

		err = con.Table("barang_kasir").Where("kode_barang_kasir = ?", kode_stock[i]).Update("jumlah", stock_baru)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
	}

	var req request_kasir.Update_Jumlah_Dan_Harga_Request

	req.Jumlah_total = tot_jumlah
	req.Total_harga = tot_harga

	err = con.Table("pembayaran").Where("kode_pembayaran = ?", Request.Kode_pembayaran).Select("jumlah_total", "total_harga").Updates(&req)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	var pembayaran response_kasir.Read_Pembayaran_Response

	pembayaran.Kode_pembayaran = Request.Kode_pembayaran
	pembayaran.Kode_nota = Request.Kode_nota

	date, _ = time.Parse("2006-01-02", Request.Tanggal)
	Request.Tanggal = date.Format("02-01-2006")
	pembayaran.Tanggal = Request.Tanggal

	pembayaran.Kode_jenis_pembayaran = Request.Kode_jenis_pembayaran

	err = con.Table("jenis_pembayaran").Select("nama_jenis_pembayaran").Where("kode_jenis_pembayaran=?", Request.Kode_jenis_pembayaran).Scan(&pembayaran.Nama_jenis_pambayaran)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("user_management").Select("nama_store").Where("kode_store=?", Request.Kode_store).Scan(&pembayaran.Nama_store)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	pembayaran.Total_harga = tot_harga
	pembayaran.Jumlah_total = tot_jumlah

	var detail_pembayaran []response_kasir.Read_Detail_Pembayaran_Response

	err = con.Table("barang_pembayaran").Select("kode_barang_pembayaran", "kode_pembayaran", "kode_barang_kasir", "nama_barang_kasir", "jumlah_barang", "nama_satuan", "harga").Where("kode_pembayaran = ?", Request.Kode_pembayaran).Scan(&detail_pembayaran)

	pembayaran.Detail_pembayaran = detail_pembayaran

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = pembayaran
	}

	return res, nil
}

func Read_Pembayaran(Request request_kasir.Read_Pembayaran_Request, Request_filter request_kasir.Read_Filter_Pembayaran_Request) (response_kasir.Response, error) {
	var res response_kasir.Response

	var arr_data []response_kasir.Read_Pembayaran_Response
	var data response_kasir.Read_Pembayaran_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT kode_pembayaran, kode_nota, DATE_FORMAT(tanggal, '%d-%m-%Y'), pembayaran.kode_jenis_pembayaran, nama_jenis_pembayaran , nama_store, total_harga, jumlah_total  FROM pembayaran join user_management um on um.kode_store = pembayaran.kode_store join jenis_pembayaran jp on jp.kode_jenis_pembayaran = pembayaran.kode_jenis_pembayaran WHERE pembayaran.kode_kasir = '" + Request.Kode_kasir + "'"

	if Request_filter.Tanggal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal)
		Request_filter.Tanggal = date.Format("2006-01-02")

		statement += " AND tanggal = '" + Request_filter.Tanggal + "'"

	}

	if Request_filter.Kode_store != "" {

		statement += " AND pembayaran.kode_store = '" + Request_filter.Kode_store + "'"

	}

	statement += " ORDER BY pembayaran.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_pembayaran, &data.Kode_nota, &data.Tanggal, &data.Kode_jenis_pembayaran, &data.Nama_jenis_pambayaran, &data.Nama_store, &data.Total_harga, &data.Jumlah_total)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_pembayaran")
		var detail_data []response_kasir.Read_Detail_Pembayaran_Response

		err = con_detail.Select("kode_barang_pembayaran", "kode_pembayaran", "kode_barang_kasir", "nama_barang_kasir", "jumlah_barang", "nama_satuan", "harga").Where("kode_pembayaran = ?", data.Kode_pembayaran).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_pembayaran = detail_data

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
