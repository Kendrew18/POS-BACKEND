package stock

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Barang(Request request.Input_Barang_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("stock")

	co := 0

	err := con.Select("co").Last(&co)

	Request.Co = co + 1
	Request.Kode_stock = "ST-" + strconv.Itoa(Request.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_stock", "nama_barang", "harga_jual", "satuan", "kode_jenis_barang").Create(&Request)

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

func Read_Barang(Request request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Barang_Response
	var detail_data []response.Read_Detail_Barang_Response

	con := db.CreateConGorm().Table("jenis_barang")

	con_barang := db.CreateConGorm().Table("stock")

	err := con.Select("kode_jenis_barang", "nama_jenis_barang").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for i := 0; i < len(data); i++ {
		err := con_barang.Select("kode_stock", "nama_barang", "harga_jual", "satuan").Where("kode_gudang = ? && kode_jenis_barang = ?", Request.Kode_gudang, data[i].Kode_jenis_barang).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data[i].Read_Detail_Barang = detail_data
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

	con := db.CreateConGorm().Table("stock")

	err := con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "satuan_barang.nama_satuan_barang", "jenis_barang.nama_jenis_barang").Joins("jenis_barang").Joins("satuan_barang").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&data).Error

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
