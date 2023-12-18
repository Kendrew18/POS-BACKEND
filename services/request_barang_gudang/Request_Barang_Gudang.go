package request_barang_gudang

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/services/stock_keluar"
	"fmt"
	"net/http"
	"time"
)

func Update_Status_Refund(Request request.Update_Status_Request_Barang_Request, Request_kode request.Update_Status_Kode_Refund_Request) (response.Response, error) {
	var res response.Response
	//var err2 error
	con := db.CreateConGorm().Table("refund")
	status := -1

	err := con.Select("status").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 3 {
		if Request.Status == 1 || (Request.Status == 4 && status == 0) {

			con := db.CreateConGorm().Table("refund")

			err := con.Where("kode_refund = ?", Request_kode.Kode_refund).Select("status").Updates(&Request)

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
		} else if Request.Status == 2 {
			var data_update request.Data_Update_Refund
			con := db.CreateConGorm().Table("refund")

			date := time.Now()
			tanggal_sekarang := date.Format("2006-01-02")

			fmt.Println(tanggal_sekarang)
			data_update.Status = Request.Status
			data_update.Tanggal_pengembalian = tanggal_sekarang

			err := con.Where("kode_refund = ?", Request_kode.Kode_refund).Select("tanggal_pengembalian", "status").Updates(&data_update)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var Request_Input_stock request.Input_Stock_Keluar_Request

			var Temp request.Tanggal_dan_Kode_Gudang

			con_user := db.CreateConGorm().Table("user")

			penanggung_jawab := ""

			err = con_user.Select("username").Where("id_user = ?", Request_kode.Kode_user).Scan(&penanggung_jawab)

			err = con.Select("tanggal_pengembalian", "kode_gudang").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&Temp)

			Request_Input_stock.Kode_gudang = Temp.Kode_gudang
			Request_Input_stock.Tanggal = Temp.Tanggal_pengembalian
			Request_Input_stock.Nama_penanggung_jawab = penanggung_jawab

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			err = con.Select("kode_supplier").Where("kode_refund = ?", Request_kode.Kode_refund).Scan(&Request_Input_stock.Kode)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			fmt.Println(Request_Input_stock.Kode)

			con_barang := db.CreateConGorm().Table("barang_refund")
			rows, err2 := con_barang.Select("kode_nota", "kode_stock", "jumlah").Where("kode_refund = ?", Request_kode.Kode_refund).Rows()

			defer rows.Close()

			if err2 != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err2
			}

			for rows.Next() {
				kode_stock := ""
				jumlah_barang := ""
				harga := ""
				var data request.Move_Refund_To_Stock_Keluar_Request
				err2 = rows.Scan(&Request_Input_stock.Kode_nota, &data.Kode_stock, &data.Jumlah_barang)
				fmt.Println(Request_Input_stock)

				if err2 != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = data
					return res, err2
				}
				kode_stock = kode_stock + "|" + data.Kode_stock + "|"

				jumlah_barang = jumlah_barang + "|" + fmt.Sprintf("%f", data.Jumlah_barang) + "|"

				harga = harga + "|" + "0" + "|"

				var Request_Barang_V2 request.Input_Barang_Stock_Keluar_Request

				Request_Barang_V2.Kode_stock = kode_stock
				Request_Barang_V2.Jumlah_barang = jumlah_barang
				Request_Barang_V2.Harga_jual = harga

				fmt.Println(Request_Barang_V2)

				res, err2 = stock_keluar.Input_Stock_Keluar(Request_Input_stock, Request_Barang_V2)
			}

			if err2 != nil {
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

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}

	return res, nil
}
