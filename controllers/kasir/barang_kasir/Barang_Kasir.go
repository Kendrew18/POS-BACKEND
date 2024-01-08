package barang_kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/barang_kasir"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputBarangKasir(c echo.Context) error {

	var Request request_kasir.Input_Barang_Kasir_Request
	Request.Nama_barang_kasir = c.FormValue("nama_barang_kasir")
	Request.Kode_satuan = c.FormValue("kode_satuan")
	Request.Jumlah_pengali = c.FormValue("jumlah_pengali")
	Request.Kode_kasir = c.FormValue("kode_kasir")
	Request.Kode_store = c.FormValue("kode_store")

	result, err := barang_kasir.Input_Barang_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadBarangKasir(c echo.Context) error {

	var Request request_kasir.Read_Barang_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := barang_kasir.Read_Barang_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
