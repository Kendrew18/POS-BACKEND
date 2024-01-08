package request_barang_kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/request_barang_kasir"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputRequestBarangKasir(c echo.Context) error {

	var Request request_kasir.Input_Request_Barang_Kasir_Request
	var Request_barang request_kasir.Input_Barang_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Kode_store = c.FormValue("kode_store")
	Request.Kode_kasir = c.FormValue("kode_kasir")
	Request.Tanggal_request = c.FormValue("tanggal_request")

	Request_barang.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request_barang.Kode_stock_gudang = c.FormValue("kode_stock_gudang")
	Request_barang.Jumlah = c.FormValue("jumlah")

	result, err := request_barang_kasir.Input_Request_Barang_Kasir(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadRequestBarangKasir(c echo.Context) error {

	var Request request_kasir.Read_Request_Barang_Kasir_Request
	var Request_filter request_kasir.Read_Filter_Request_Barang_Kasir

	Request.Kode_kasir = c.FormValue("kode_kasir")

	Request_filter.Kode_store = c.FormValue("kode_store")
	Request_filter.Tanggal_1 = c.FormValue("tanggal")

	result, err := request_barang_kasir.Read_Request_Barang_Kasir(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateRequestBarangKasir(c echo.Context) error {

	var Request request_kasir.Update_Request_Barang_Kasir_Request
	var Request_kode request_kasir.Update_Request_Barang_Kasir_Kode

	Request.Kode_stock_gudang = c.FormValue("kode_stock_gudang")
	Request.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request.Jumlah, _ = strconv.ParseFloat(c.FormValue("jumlah"), 64)

	Request_kode.Kode_barang_request = c.FormValue("kode_barang_request")

	result, err := request_barang_kasir.Update_Request_Barang_Kasir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteRequestBarangKasir(c echo.Context) error {
	var Request_kode request_kasir.Update_Request_Barang_Kasir_Kode

	Request_kode.Kode_barang_request = c.FormValue("kode_barang_request")

	result, err := request_barang_kasir.Delete_Request_Barang_Kasir(Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusRequestBarangKasir(c echo.Context) error {
	var Request request_kasir.Update_Status_Request_Barang_Kasir
	var Request_kode request_kasir.Kode_Request_Barang_Kasir_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	Request_kode.Kode_request_barang_kasir = c.FormValue("kode_request_barang_kasir")

	result, err := request_barang_kasir.Update_Status_Request_Barang_Kasir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func Dropdownstatus(c echo.Context) error {
	var Request request_kasir.Dropdown_Status_Kasir_Request

	Request.Kode = c.FormValue("kode")
	Request.Kode_request_barang_kasir = c.FormValue("kode_request_barang_kasir")

	result, err := request_barang_kasir.Dropdown_status(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
