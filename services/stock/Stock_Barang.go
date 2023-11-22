package stock

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Input_Barang(Request request.Input_Barang_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("stock")

	nama_barang := tools.String_Separator_To_String(Request.Nama_Barang)
	harga_jual := tools.String_Separator_To_Int64(Request.Harga_jual)
	kode_satuan_barang := tools.String_Separator_To_String(Request.Kode_satuan_barang)

	cont := int64(0)

	for i := 0; i < len(nama_barang); i++ {
		var Request_V2 request.Input_Barang_Request_V2
		con_check := db.CreateConGorm()

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_V2.Co = co + 1
		Request_V2.Kode_stock = "ST-" + strconv.Itoa(Request_V2.Co)
		Request_V2.Nama_Barang = nama_barang[i]
		Request_V2.Harga_jual = harga_jual[i]
		Request_V2.Kode_satuan_barang = kode_satuan_barang[i]
		Request_V2.Kode_jenis_barang = Request.Kode_jenis_barang
		Request_V2.Kode_gudang = Request.Kode_gudang

		fmt.Println(co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = co
			return res, err.Error
		}

		nb := strings.ToLower(Request.Nama_Barang)

		check := ""

		err = con_check.Raw("SELECT LOWER(nama_barang) FROM stock WHERE nama_barang=@nama_barang && kode_gudang=@kode_gudang", sql.Named("nama_barang", nb), sql.Named("kode_gudang", Request_V2.Kode_gudang)).Scan(&check)

		fmt.Println(check)

		if check == "" {
			err = con.Select("co", "kode_stock", "nama_barang", "harga_jual", "kode_satuan_barang", "kode_jenis_barang", "kode_gudang").Create(&Request_V2)
		}

		cont = err.RowsAffected

	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": cont,
	}

	return res, nil
}

func Read_Barang(Request request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Barang_Response
	var obj_data response.Read_Barang_Response

	con := db.CreateConGorm().Table("jenis_barang")

	rows, err := con.Select("kode_jenis_barang", "nama_jenis_barang").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		con_barang := db.CreateConGorm().Table("stock")
		var detail_data []response.Read_Detail_Barang_Response
		rows.Scan(&obj_data.Kode_jenis_barang, &obj_data.Jenis_Barang)

		err := con_barang.Select("kode_stock", "nama_barang", "harga_jual", "satuan_barang.nama_satuan_barang").Joins("join satuan_barang on satuan_barang.kode_satuan_barang = stock.kode_satuan_barang").Where("stock.kode_gudang = ? && kode_jenis_barang = ?", Request.Kode_gudang, obj_data.Kode_jenis_barang).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		obj_data.Read_Detail_Barang = detail_data

		data = append(data, obj_data)
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}

	return res, nil
}

func Delete_Barang(Request request.Delete_Barang_Request) (response.Response, error) {

	var res response.Response
	var barang_stock_masuk []string
	var barang_stock_keluar []string
	var barang_supplier []string

	con_masuk := db.CreateConGorm().Table("barang_stock_masuk")

	err := con_masuk.Select("kode_stock").Where("kode_stock =?", Request.Kode_stock).Scan(&barang_stock_masuk).Error

	con_keluar := db.CreateConGorm().Table("barang_stock_keluar")

	err = con_keluar.Select("kode_stock").Where("kode_stock =?", Request.Kode_stock).Scan(&barang_stock_keluar).Error

	con_supplier := db.CreateConGorm().Table("barang_supplier")

	err = con_supplier.Select("kode_stock").Where("kode_stock =?", Request.Kode_stock).Scan(&barang_supplier).Error

	if barang_stock_masuk == nil && barang_stock_keluar == nil && barang_supplier == nil && err == nil {
		con := db.CreateConGorm().Table("stock")

		err := con.Where("kode_stock=?", Request.Kode_stock).Delete("")

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
		res.Message = "Erorr karena ada condition yang tidak terpenuhi"
		res.Data = Request
		return res, err
	}

	return res, nil
}

func Read_Stock(Request request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var data response.Read_Stock_Response
	var arr_data []response.Read_Stock_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm().Table("stock")

	if Request.Kode_jenis_barang == "" {
		rows, err = con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "sb.nama_satuan_barang", "jb.nama_jenis_barang").Joins("JOIN jenis_barang jb ON jb.kode_jenis_barang = stock.kode_jenis_barang").Joins("JOIN satuan_barang sb ON sb.kode_satuan_barang = stock.kode_satuan_barang").Where("stock.kode_gudang = ?", Request.Kode_gudang).Order("stock.co ASC").Rows()
	} else {
		rows, err = con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "sb.nama_satuan_barang", "jb.nama_jenis_barang").Joins("JOIN jenis_barang jb ON jb.kode_jenis_barang = stock.kode_jenis_barang").Joins("JOIN satuan_barang sb ON sb.kode_satuan_barang = stock.kode_satuan_barang").Where("stock.kode_gudang = ? && stock.kode_jenis_barang = ?", Request.Kode_gudang, Request.Kode_jenis_barang).Order("stock.co ASC").Rows()
	}

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Kode_stock, &data.Nama_Barang, &data.Harga_jual, &data.Jumlah, &data.Nama_satuan_barang, &data.Nama_jenis_barang)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_gudang := db.CreateConGorm().Table("gudang")

		status := 0

		err = con_gudang.Select("status_lifo_fifo").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&status).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status tidak ada"
			res.Data = data
			return res, err
		}

		var data_detail []response.Detail_Stock_Response

		con_stock_masuk := db.CreateConGorm().Table("stock_keluar_masuk")

		if status == 0 {
			err = con_stock_masuk.Select("DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "jumlah_barang", "harga").Joins("JOIN detail_stock bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("kode_stock=? && stock_keluar_masuk.status=0", data.Kode_stock).Order("tanggal DESC").Scan(&data_detail).Error
		} else {
			err = con_stock_masuk.Select("DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "jumlah_barang", "harga").Joins("JOIN detail_stock bs ON bs.kode_stock_keluar_masuk = stock_keluar_masuk.kode_stock_keluar_masuk").Where("kode_stock=? && stock_keluar_masuk.status=0", data.Kode_stock).Order("tanggal ASC").Scan(&data_detail).Error
		}

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_stock = data_detail

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

func Detail_stock(Request request.Read_Detail_Stock) (response.Response, error) {
	var res response.Response

	var data response.Read_Detail_Stock_Response
	var arr_data []response.Read_Detail_Stock_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	con_gudang := db.CreateConGorm().Table("gudang")

	status := 0

	err = con_gudang.Select("status_lifo_fifo").Where("kode_gudang=?", Request.Kode_gudang).Scan(&status).Error

	if status == 0 {
		rows, err = con.Raw("SELECT kode_barang_keluar_masuk, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, if(status=0,'Masuk','Keluar') AS keterangan, nama_barang, bs.jumlah_barang FROM stock_keluar_masuk AS skm JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = skm.kode_stock_keluar_masuk JOIN stock ON bs.kode_stock=stock.kode_stock WHERE bs.kode_stock=? ORDER BY tanggal DESC", Request.Kode_stock).Rows()
	} else {
		rows, err = con.Raw("SELECT kode_barang_keluar_masuk, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, if(status=0,'Masuk','Keluar') AS keterangan, nama_barang, bs.jumlah_barang FROM stock_keluar_masuk AS skm JOIN barang_stock_keluar_masuk bs ON bs.kode_stock_keluar_masuk = skm.kode_stock_keluar_masuk JOIN stock ON bs.kode_stock=stock.kode_stock WHERE bs.kode_stock=? ORDER BY tanggal ASC", Request.Kode_stock).Rows()
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Kode_barang_keluar_masuk, &data.Tanggal, &data.Keterangan, &data.Nama_barang, &data.Jumlah)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

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

func Dropdown_Stock(Request request.Dropdown_Stock_Request) (response.Response, error) {
	var res response.Response
	var data []response.Dropdown_Nama_Barang_Response

	con := db.CreateConGorm().Table("stock")

	err := con.Select("kode_stock", "nama_barang").Where("kode_gudang = ? && jumlah > 0", Request.Kode_gudang).Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}
	return res, nil
}
