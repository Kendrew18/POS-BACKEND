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

func Read_Stock(Request request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Stock_Response
	var err error

	con := db.CreateConGorm().Table("stock")

	if Request.Kode_jenis_barang != "" {
		err = con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "satuan_barang.nama_satuan_barang", "jenis_barang.nama_jenis_barang").Joins("jenis_barang").Joins("satuan_barang").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&data).Error
	} else {
		err = con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "satuan_barang.nama_satuan_barang", "jenis_barang.nama_jenis_barang").Joins("jenis_barang").Joins("satuan_barang").Where("kode_gudang = ? && kode_jenis_barang = ?", Request.Kode_gudang, Request.Kode_jenis_barang).Scan(&data).Error
	}

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

func Dropdown_Nama_Barang(Request request.Dropdown_Nama_Barang_Request) (response.Response, error) {

	var res response.Response
	var Nama_barang []response.Read_Dropdown_Nama_Supplier_Response

	con := db.CreateConGorm().Table("barang_supplier")

	err := con.Select("kode_stock", "nama_barang").Joins("JOIN stock on stock.kode_stock = barang_supplier.kode_stock").Where("kode_supplier = ?", Request.Kode_supplier).Scan(&Nama_barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_barang
		return res, err
	}

	if Nama_barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Nama_barang
	}

	return res, nil
}
