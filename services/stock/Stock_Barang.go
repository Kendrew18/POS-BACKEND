package stock

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Barang(Barang request.Input_Barang_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("stock")

	co := 0

	err := con.Select("co").Last(&co)

	Barang.Co = co + 1
	Barang.Kode_stock = "ST-" + strconv.Itoa(Barang.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_stock", "nama_barang", "harga_jual", "satuan", "kode_jenis_barang").Create(&Barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang
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

func Read_Barang(kode_gudang request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var Barang []response.Read_Barang_Response
	var Detail_Barang []response.Read_Detail_Barang_Response

	con := db.CreateConGorm().Table("jenis_barang")

	con_barang := db.CreateConGorm().Table("stock")

	err := con.Select("kode_jenis_barang", "nama_jenis_barang").Where("kode_gudang = ?", kode_gudang.Kode_gudang).Scan(&Barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang
		return res, err
	}

	for i := 0; i < len(Barang); i++ {
		err := con_barang.Select("kode_stock", "nama_barang", "harga_jual", "satuan").Where("kode_gudang = ? && kode_jenis_barang = ?", kode_gudang.Kode_gudang, Barang[i].Kode_jenis_barang).Scan(&Detail_Barang).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Barang
			return res, err
		}

		Barang[i].Read_Detail_Barang = Detail_Barang
	}

	if Barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Barang
	}

	return res, nil
}

func Read_Stock(kode_gudang request.Read_Stock_Request) (response.Response, error) {

	var res response.Response
	var Barang []response.Read_Stock_Response

	con := db.CreateConGorm().Table("stock")

	err := con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "satuan", "jenis_barang.nama_jenis_barang").Joins("jenis_barang").Where("kode_gudang = ?", kode_gudang.Kode_gudang).Scan(&Barang).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang
		return res, err
	}

	if Barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Barang
	}

	return res, nil
}
