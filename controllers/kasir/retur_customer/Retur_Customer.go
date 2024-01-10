package retur_customer

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/retur_customer"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputReturCustomer(c echo.Context) error {

	var Request request_kasir.Input_Retur_Customer_Request
	var Request_Barang request_kasir.Input_Barang_Retur_Customer_Request

	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Tanggal = c.FormValue("tanggal")
	Request.Kode_bentuk_retur = c.FormValue("kode_bentuk_retur")
	Request.Kode_store = c.FormValue("kode_store")
	Request.Kode_kasir = c.FormValue("kode_kasir")

	Request_Barang.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request_Barang.Jumlah = c.FormValue("jumlah")
	Request_Barang.Keterangan = c.FormValue("keterangan")

	result, err := retur_customer.Input_Retur_Customer(Request, Request_Barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadReturCustomer(c echo.Context) error {

	var Request request_kasir.Read_Retur_Customer_Request
	var Request_Filter request_kasir.Read_Filter_Retur_Customer_Request

	Request.Kode_kasir = c.FormValue("kode_kasir")

	Request_Filter.Tanggal = c.FormValue("tanggal")
	Request_Filter.Kode_store = c.FormValue("kode_store")

	result, err := retur_customer.Read_Retur_Customer(Request, Request_Filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DropdownKodeNotaReturCustomer(c echo.Context) error {

	var Request request_kasir.Read_Dropdown_Kode_Nota_Request

	Request.Kode_kasir = c.FormValue("kode_kasir")

	Request.Tanggal = c.FormValue("tanggal")

	result, err := retur_customer.Dropdown_Kode_Nota_Retur_Customer(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateReturCustomer(c echo.Context) error {

	var Request request_kasir.Update_Retur_Customer_Request
	var Request_kode request_kasir.Update_Kode_Retur_Customer_Request

	Request.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request.Jumlah, _ = strconv.ParseFloat(c.FormValue("jumlah"), 64)
	Request.Keterangan = c.FormValue("keterangan")

	Request_kode.Kode_barang_retur_customer = c.FormValue("kode_barang_retur_customer")

	result, err := retur_customer.Update_Retur_Customer(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteRequestBarangKasir(c echo.Context) error {

	var Request request_kasir.Update_Kode_Retur_Customer_Request

	Request.Kode_barang_retur_customer = c.FormValue("kode_barang_retur_customer")

	result, err := retur_customer.Delete_Request_Barang_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusRequestBarangKasir(c echo.Context) error {
	var Request request_kasir.Update_Status_Retur_Customer_Request
	var Request_kode request_kasir.Update_Status_Retur_Customer_Kode_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	Request_kode.Kode_retur_customer = c.FormValue("kode_retur_customer")

	result, err := retur_customer.Update_Status_Request_Barang_Kasir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DropdownBarangKasirRetur(c echo.Context) error {
	var Request request_kasir.Read_Dropdown_Barang_Retur_Request

	Request.Kode_nota = c.FormValue("kode_nota")

	result, err := retur_customer.Dropdown_Barang_Kasir_Retur(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
