package pembayaran

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/pembayaran"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputPembayaran(c echo.Context) error {
	var Request request_kasir.Input_Pembayaran_Request
	var Request_barang request_kasir.Barang_Input_Pembayaran_Request

	Request.Kode_kasir = c.FormValue("kode_kasir")
	Request.Kode_store = c.FormValue("kode_store")
	Request.Kode_jenis_pembayaran = c.FormValue("kode_jenis_pembayaran")
	Request.Tanggal = c.FormValue("tanggal")

	Request_barang.Harga = c.FormValue("harga")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request_barang.Nama_satuan = c.FormValue("nama_satuan")

	result, err := pembayaran.Input_Pembayaran(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadPembayaran(c echo.Context) error {
	var Request request_kasir.Read_Pembayaran_Request
	var Request_filter request_kasir.Read_Filter_Pembayaran_Request

	Request.Kode_kasir = c.FormValue("kode_kasir")

	Request_filter.Tanggal = c.FormValue("tanggal")
	Request_filter.Kode_store = c.FormValue("kode_store")

	result, err := pembayaran.Read_Pembayaran(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
