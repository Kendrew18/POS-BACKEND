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

func Input_Audit_Stock(Request request.Input_Audit_stock_Request, Request_detail request.Input_Detail_Audit_stock_Request, Request_user request.Input_Audit_stock_User_Request) (response.Response, error) {
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

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	tanggal_masuk := tools.String_Separator_To_String(Request_detail.Tanggal_masuk)
	stock_dalam_sistem := tools.String_Separator_To_float64(Request_detail.Stock_dalam_sistem)
	stock_rill := tools.String_Separator_To_float64(Request_detail.Stock_rill)
	selisih_stock := tools.String_Separator_To_float64(Request_detail.Selisih_stock)
	kode_bkm := tools.String_Separator_To_String(Request_detail.Kode_barang_keluar_masuk)

	total_stock := float64(0.0)

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
		fmt.Println(kode_bkm[i])

		con_bkm := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		kode_sup := ""

		err = con_bkm.Select("kode").Joins("JOIN stock_keluar_masuk ON stock_keluar_masuk.kode_stock_keluar_masuk = barang_stock_keluar_masuk.kode_stock_keluar_masuk ").Where("kode_barang_keluar_masuk = ?", kode_bkm[i]).Scan(&kode_sup)

		detail.Kode_supplier = kode_sup

		err = con_detail.Select("co", "kode_detail_audit", "kode_audit", "tanggal_masuk", "stock_dalam_sistem", "stock_rill", "selisih_stock", "kode_supplier").Create(&detail)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = detail
			return res, err.Error
		}

		con_detail_stock := db.CreateConGorm().Table("detail_stock")

		err = con_detail_stock.Where("kode_barang_keluar_masuk = ?", kode_bkm[i]).Update("jumlah_barang", &stock_rill[i])

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = detail
			return res, err.Error
		}

		total_stock = total_stock + stock_rill[i]

	}

	con_stock := db.CreateConGorm().Table("stock")

	jumlah_stock := 0.0

	err = con_stock.Select("jumlah").Where("kode_stock = ? ").Scan(&jumlah_stock)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	stock_stat := total_stock - jumlah_stock

	if stock_stat < 0.0 {
		//Keluar

		var Request_keluar request.Input_Stock_Keluar_Request
		con_stock_keluar := db.CreateConGorm().Table("stock_keluar_masuk")

		co := 0

		err := con_stock_keluar.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_keluar.Co = co + 1
		Request_keluar.Kode_stock_keluar_masuk = Request.Kode_audit
		Request_keluar.Status = 3

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		date, _ := time.Parse("02-01-2006", Request.Tanggal)
		Request_keluar.Tanggal = date.Format("2006-01-02")
		Request_keluar.Kode_nota = ""

		con_user := db.CreateConGorm().Table("user")

		username := ""
		err = con_user.Select("username").Where("id_user = ?", Request_user.Kode_user).Scan(&username)

		Request_keluar.Nama_penanggung_jawab = username
		Request_keluar.Kode_gudang = Request.Kode_gudang
		Request_keluar.Kode = ""

		err = con_stock_keluar.Select("co", "kode_stock_keluar_masuk", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode", "kode_gudang", "status").Create(&Request)

		var barang_V2 request.Input_Barang_Stock_Keluar_V2_Request

		con_bsm := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		co = 0

		err = con_bsm.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar_masuk = "BKM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar_masuk = Request_keluar.Kode_stock_keluar_masuk
		barang_V2.Kode_stock = Request.Kode_stock
		barang_V2.Jumlah_barang = jumlah_stock - total_stock
		barang_V2.Harga = 0
		barang_V2.Total_harga = 0

		fmt.Println(barang_V2)

		err = con_bsm.Select("co", "kode_barang_keluar_masuk", "kode_stock_keluar_masuk", "kode_stock", "jumlah_barang", "harga", "total_harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

	} else if stock_stat == 0.0 {

		var Request_stock_masuk request.Input_Stock_Masuk_Request

		con := db.CreateConGorm().Table("stock_keluar_masuk")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_stock_masuk.Co = co + 1
		Request_stock_masuk.Kode_stock_keluar_masuk = Request.Kode_audit

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_stock_masuk
			return res, err.Error
		}

		date, _ := time.Parse("02-01-2006", Request.Tanggal)
		Request_stock_masuk.Tanggal = date.Format("2006-01-02")
		Request_stock_masuk.Status = 4
		Request_stock_masuk.Kode_nota = ""
		Request_stock_masuk.Kode = ""
		Request_stock_masuk.Kode_gudang = Request.Kode_gudang

		con_user := db.CreateConGorm().Table("user")

		username := ""
		err = con_user.Select("username").Where("id_user = ?", Request_user.Kode_user).Scan(&username)
		Request_stock_masuk.Nama_penanggung_jawab = username

		err = con.Select("co", "kode_stock_keluar_masuk", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode", "kode_gudang", "status").Create(&Request_stock_masuk)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_stock_masuk
			return res, err.Error
		}

		var barang_V2 request.Input_Barang_Stock_Masuk_V2_Request

		con_barang := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		co = 0

		err = con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar_masuk = "BKM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar_masuk = Request_stock_masuk.Kode_stock_keluar_masuk
		barang_V2.Kode_stock = Request.Kode_stock
		barang_V2.Jumlah_barang = jumlah_stock - total_stock
		barang_V2.Harga = 0
		barang_V2.Total_harga = 0

		date3, _ := time.Parse("02-01-2006", "01-01-0001")
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

	} else if stock_stat > 0.0 {
		//Masuk

		var Request_stock_masuk request.Input_Stock_Masuk_Request

		con := db.CreateConGorm().Table("stock_keluar_masuk")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_stock_masuk.Co = co + 1
		Request_stock_masuk.Kode_stock_keluar_masuk = Request.Kode_audit

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_stock_masuk
			return res, err.Error
		}

		date, _ := time.Parse("02-01-2006", Request.Tanggal)
		Request_stock_masuk.Tanggal = date.Format("2006-01-02")
		Request_stock_masuk.Status = 2
		Request_stock_masuk.Kode_nota = ""
		Request_stock_masuk.Kode = ""
		Request_stock_masuk.Kode_gudang = Request.Kode_gudang

		con_user := db.CreateConGorm().Table("user")

		username := ""
		err = con_user.Select("username").Where("id_user = ?", Request_user.Kode_user).Scan(&username)
		Request_stock_masuk.Nama_penanggung_jawab = username

		err = con.Select("co", "kode_stock_keluar_masuk", "tanggal", "kode_nota", "nama_penanggung_jawab", "kode", "kode_gudang", "status").Create(&Request_stock_masuk)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request_stock_masuk
			return res, err.Error
		}

		var barang_V2 request.Input_Barang_Stock_Masuk_V2_Request

		con_barang := db.CreateConGorm().Table("barang_stock_keluar_masuk")

		co = 0

		err = con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar_masuk = "BKM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar_masuk = Request_stock_masuk.Kode_stock_keluar_masuk
		barang_V2.Kode_stock = Request.Kode_stock
		barang_V2.Jumlah_barang = jumlah_stock - total_stock
		barang_V2.Harga = 0
		barang_V2.Total_harga = 0

		date3, _ := time.Parse("02-01-2006", "01-01-0001")
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

	}

	err = con_stock.Where("kode_stock = ?", Request.Kode_stock).Update("jumlah", &total_stock)

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

	rows, err := con.Select("audit.kode_audit", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "nama_barang", "SUM(ds.stock_dalam_sistem)", "SUM(stock_rill)", "SUM(selisih_stock)").Joins("JOIN stock s ON s.kode_stock = audit.kode_stock").Joins("JOIN detail_audit ds ON ds.kode_audit = audit.kode_audit").Where("audit.kode_gudang = ? && tanggal = ?", Request.Kode_gudang, Request.Tanggal).Group("audit.kode_audit").Order("audit.co DESC").Rows()

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
