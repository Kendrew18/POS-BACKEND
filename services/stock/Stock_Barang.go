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

func Read_Stock(kode_gudang request.Read_Jenis_Barang_Request) (response.Response, error) {

	var res response.Response
	var Barang []response.Read_Stock_Response

	con := db.CreateConGorm().Table("stock")

	err := con.Select("kode_stock", "nama_barang", "harga_jual", "jumlah", "satuan", "jenis_barang.nama_jenis_barang").Joins("jenis_barang").Where("kode_gudang", kode_gudang.Kode_gudang).Scan(&Barang).Error

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
