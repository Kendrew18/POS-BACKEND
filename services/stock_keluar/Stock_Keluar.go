package stock_keluar

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

func Input_Stock_Keluar(Request request.Input_Stock_Keluar_Request, Request_Barang request.Input_Barang_Stock_Keluar_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("stock")

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	jumlah := tools.String_Separator_To_float64(Request_Barang.Jumlah_barang)

	for i := 0; i < len(kode_stock); i++ {
		nama_barang := ""
		err := con.Select("nama_barang").Where("kode_stock = ? && jumlah >= ?", kode_stock[i], jumlah[i]).Scan(&nama_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = nama_barang + " out of stock"
			res.Data = kode_stock[i]
			return res, err.Error
		}
	}

	con_stock_keluar := db.CreateConGorm().Table("stock_keluar")

	co := 0

	err := con_stock_keluar.Select("co").Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_stock_keluar = "SK-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal_keluar)
	Request.Tanggal_keluar = date.Format("2006-01-02")

	err = con_stock_keluar.Select("co", "kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "kode_toko", "kode_gudang").Create(&Request)

	harga_jual := tools.String_Separator_To_Int64(Request_Barang.Harga_jual)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Stock_Keluar_V2_Request

		nama_barang := ""

		err = con.Select("nama_barang").Where("kode_stock = ? && jumlah >= ?", kode_stock[i], jumlah[i]).Scan(&nama_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = nama_barang + " out of stock"
			res.Data = kode_stock[i]
			return res, err.Error
		}

		con_bsm := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		co := 0

		err := con_bsm.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar_masuk = "BSM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar_masuk = Request.Kode_stock_keluar
		barang_V2.Kode_stock = kode_stock[i]
		barang_V2.Jumlah_barang = jumlah[i]
		barang_V2.Harga = harga_jual[i]
		barang_V2.Total_harga = int64(math.Round(float64(harga_jual[i]) * jumlah[i]))
		barang_V2.Status = 1

		fmt.Println(barang_V2)

		err = con_bsm.Select("co", "kode_barang_keluar_masuk", "kode_stock_keluar_masuk", "kode_stock", "jumlah_barang", "harga", "total_harga", "status").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		jumlah_lama := 0.0

		con_stock := db.CreateConGorm().Table("stock")

		err = con_stock.Select("jumlah").Where("kode_stock=?", barang_V2.Kode_stock).Scan(&jumlah_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "not found"
			res.Data = kode_stock[i]
			return res, err.Error
		}

		Jumlah_baru := jumlah_lama - barang_V2.Jumlah_barang

		err = con_stock.Where("kode_stock = ?", barang_V2.Kode_stock).Update("jumlah", Jumlah_baru)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update gagal"
			res.Data = barang_V2
			return res, err.Error
		}

		con_gudang := db.CreateConGorm().Table("gudang")

		status := 0

		err = con_gudang.Select("status_lifo_fifo").Where("kode_gudang=?", Request.Kode_gudang).Scan(&status)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "read gudang gagal"
			res.Data = barang_V2
			return res, err.Error
		}

		var data []response.Read_Tanggal_dan_Jumlah

		con_stock_masuk := db.CreateConGorm().Table("stock_masuk")
		//0 lifo
		//1 fifo
		if status == 0 {

			err = con_stock_masuk.Select("kode_barang_keluar_masuk", "kode_stock", "stock_masuk.tanggal_masuk", "b.jumlah_barang").Joins("JOIN barang_stock_keluar_masuk b on b.kode_stock_keluar_masuk = stock_masuk.kode_stock_masuk").Where("kode_stock = ?", kode_stock[i]).Order("tanggal_masuk DESC").Scan(&data)
		} else {
			err = con_stock_masuk.Select("kode_barang_keluar_masuk", "kode_stock", "stock_masuk.tanggal_masuk", "b.jumlah_barang").Joins("JOIN barang_stock_keluar_masuk b on b.kode_stock_keluar_masuk = stock_masuk.kode_stock_masuk").Where("kode_stock = ?", kode_stock[i]).Order("tanggal_masuk ASC").Scan(&data)
		}

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "read stock masuk gagal"
			res.Data = barang_V2
			return res, err.Error
		}
		//fmt.Println(data)

		x := 0

		index := 0

		jmlh := jumlah[i]

		update_jumlah := 0.0

		//fmt.Println(data)

		for x == 0 && index < len(data) && data != nil {

			con_detail_stock := db.CreateConGorm().Table("detail_stock")

			if data[index].Jumlah_barang >= jmlh {
				jmlh = data[index].Jumlah_barang - jmlh
				x = 1
				update_jumlah = jmlh
			} else {
				jmlh = jmlh - data[index].Jumlah_barang
				update_jumlah = 0.0
			}

			err = con_detail_stock.Where("kode_barang_keluar_masuk = ?", data[index].Kode_barang_keluar_masuk).Update("jumlah_barang", update_jumlah)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Update gagal"
				res.Data = barang_V2
				return res, err.Error
			}

			index++
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

func Read_Stock_Keluar(Request request.Read_Stock_Keluar_Request, Request_filter request.Filter_Stock_Keluar_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Stock_Keluar_Response
	var data response.Read_Stock_Keluar_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm().Table("stock_keluar")

	if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" && Request_filter.Kode_toko != "" {
		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_2)
		date_sql2 := date2.Format("2006-01-02")

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar").Where("stock_keluar.kode_gudang = ? && (tanggal_masuk >= ? && tanggal_masuk <= ?) && stock_keluar.kode_toko = ?", Request.Kode_gudang, date_sql, date_sql2, Request_filter.Kode_toko).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 == "" && Request_filter.Kode_toko != "" {
		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar ").Where("stock_keluar.kode_gudang = ? && tanggal_keluar = ? && stock_keluar.kode_toko = ?", Request.Kode_gudang, date_sql, Request_filter.Kode_toko).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 != "" && Request_filter.Kode_toko == "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_2)
		date_sql2 := date2.Format("2006-01-02")

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar ").Where("stock_keluar.kode_gudang = ? && (tanggal_keluar >= ? && tanggal_keluar <= ?)", Request.Kode_gudang, date_sql, date_sql2).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	} else if Request_filter.Tanggal_1 != "" && Request_filter.Tanggal_2 == "" && Request_filter.Kode_toko == "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_1)
		date_sql := date.Format("2006-01-02")

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar ").Where("stock_keluar.kode_gudang = ? && tanggal_keluar = ?", Request.Kode_gudang, date_sql).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	} else if Request_filter.Kode_toko != "" {

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar ").Where("stock_keluar.kode_gudang = ? && stock_keluar.kode_toko = ?", Request.Kode_gudang, Request_filter.Kode_toko).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	} else {

		rows, err = con.Select("stock_keluar.kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "s.nama_supplier", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = stock_keluar.kode_stock_keluar ").Where("stock_keluar.kode_gudang = ? ", Request.Kode_gudang).Group("stock_keluar.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	}

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&data.Kode_stock_keluar, &data.Tanggal_keluar, &data.Kode_nota, &data.Penanggung_jawab, &data.Nama_toko, &data.Jumlah_total, &data.Total_harga)
		con_detail := db.CreateConGorm().Table("barang_stock_keluar_masuk")
		var detail_data []response.Read_Barang_Stock_Keluar_Response

		err := con_detail.Select("kode_barang_keluar_masuk", "nama_barang", "tanggal_kadaluarsa", "jumlah_barang", "harga").Joins("join stock s on barang_stock_keluar_masuk.kode_stock = s.kode_stock").Where("kode_stock_keluar_masuk = ?", data.Kode_stock_keluar).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Read_Barang_Stock_Keluar_Request = detail_data

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
