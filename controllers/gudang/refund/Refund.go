package refund

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/gudang/refund_supplier"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputRefundSupplier(c echo.Context) error {

	var Request request.Input_Refund_Request
	var Request_Barang request.Input_Barang_Refund_Request

	Request.Tanggal = c.FormValue("tanggal")
	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_Barang.Jumlah = c.FormValue("jumlah")
	Request_Barang.Kode_stock = c.FormValue("kode_stock")
	Request_Barang.Kode_nota = c.FormValue("kode_nota")
	Request_Barang.Keterangan = c.FormValue("keterangan")
	Request_Barang.Tanggal_stock_masuk = c.FormValue("tanggal_stock_masuk")

	result, err := refund_supplier.Input_Refund_Supplier(Request, Request_Barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadRefund(c echo.Context) error {
	var Request request.Read_Refund_Request
	var Request_Filter request.Read_Refund_Filter_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_Filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_Filter.Tanggal_2 = c.FormValue("tanggal_2")
	Request_Filter.Kode_supplier = c.FormValue("kode_supplier")

	result, err := refund_supplier.Read_Refund(Request, Request_Filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateBarangRefund(c echo.Context) error {
	var Request request.Update_Refund_Request
	var Request_barang request.Update_Barang_Refund_Request
	Request.Kode_barang_refund = c.FormValue("kode_barang_refund")

	Request_barang.Kode_nota = c.FormValue("kode_nota")
	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Tanggal_stock_masuk = c.FormValue("tanggal_stock_masuk")
	Request_barang.Jumlah, _ = strconv.ParseFloat(c.FormValue("jumlah"), 64)
	Request_barang.Keterangan = c.FormValue("keterangan")

	result, err := refund_supplier.Update_Barang_Refund(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteBarangRefund(c echo.Context) error {
	var Request request.Update_Refund_Request
	Request.Kode_barang_refund = c.FormValue("kode_barang_refund")

	result, err := refund_supplier.Delete_Barang_Refund(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusRefund(c echo.Context) error {
	var Request request.Update_Status_Refund_Request
	var Request_kode request.Update_Status_Kode_Refund_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	Request_kode.Kode_refund = c.FormValue("kode_refund")
	Request_kode.Kode_user = c.FormValue("kode_user")

	result, err := refund_supplier.Update_Status_Refund(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
