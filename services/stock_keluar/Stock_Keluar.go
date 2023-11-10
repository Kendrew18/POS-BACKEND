package stock_keluar

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
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

	err := con.Select("co").Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_stock_keluar = "SM-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("2006-01-02", Request.Tanggal_keluar)

	Request.Tanggal_keluar = date.Format("2006-01-02")

	err = con_stock_keluar.Select("co", "kode_stock_keluar", "kode_gudang", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "kode_toko").Create(&Request)

	harga_jual := tools.String_Separator_To_Int64(Request_Barang.Harga_jual)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Stock_Keluar_V2_Request

		con := db.CreateConGorm().Table("barang_stock_keluar")

		co := 0

		err := con.Select("co").Order("co DESC").Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_keluar = "BSM-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_stock_keluar = Request.Kode_stock_keluar
		barang_V2.Kode_stock = kode_stock[i]
		barang_V2.Jumlah_barang = jumlah[i]
		barang_V2.Harga_jual = harga_jual[i]
		barang_V2.Total_harga = int64(math.Round(float64(harga_jual[i]) * jumlah[i]))

		err = con.Select("co", "kode_barang_keluar", "kode_stock_keluar", "kode_stock", "jumlah_barang", "harga_jual", "total_harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		con_stock := db.CreateConGorm().Table("stock")

		nama_barang := ""

		err = con.Select("nama_barang").Where("kode_stock = ? && jumlah >= ?", kode_stock[i], jumlah[i]).Scan(&nama_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = nama_barang + " out of stock"
			res.Data = kode_stock[i]
			return res, err.Error
		}
		jumlah_lama := 0.0
		err = con_stock.Select("jumlah").Where("kode_stock=?", barang_V2.Kode_stock).Scan(&jumlah_lama)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "not found"
			res.Data = kode_stock[i]
			return res, err.Error
		}

		Jumlah_baru := jumlah_lama - barang_V2.Jumlah_barang

		err = con_stock.Where("kode_stock = ?", barang_V2.Kode_stock).Update("jumlah_barang", Jumlah_baru)
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

func Read_Stock_Keluar(Request request.Read_Stock_Keluar_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Stock_Keluar_Response
	var data response.Read_Stock_Keluar_Response

	con := db.CreateConGorm().Table("stock_masuk")

	rows, err := con.Select("kode_stock_keluar", "tanggal_keluar", "kode_nota", "nama_penanggung_jawab", "t.nama_toko", "sum(jumlah_barang)", "sum(total_harga)").Joins("JOIN toko t ON t.kode_toko = stock_keluar.kode_toko").Joins("JOIN barang_stock_keluar bs ON bs.kode_stock_keluar = stock_keluar.kode_stock_keluar ").Where("kode_gudang = ?", Request.Kode_gudang).Group("bs.kode_stock_keluar").Order("stock_keluar.co DESC").Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&data.Kode_stock_keluar, &data.Tanggal_keluar, &data.Kode_nota, &data.Nama_penanggung_jawab, &data.Nama_toko)
		con_detail := db.CreateConGorm().Table("barang_stock_masuk")
		var detail_data []response.Read_Barang_Stock_Keluar_Response

		err := con_detail.Select("kode_barang_keluar", "nama_barang", "jumlah_barang", "harga_jual").Joins("join stock s on barang_stock_masuk.kode_stock = s.kode_stock").Where("kode_stock_keluar = ?", data.Kode_stock_keluar).Scan(&detail_data).Error

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
