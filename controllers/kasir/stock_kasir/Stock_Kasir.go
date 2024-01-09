package stock_kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/stock_kasir"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ReadStockKasir(c echo.Context) error {

	var Request request_kasir.Read_Stock_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")
	Request.Kode_store = c.FormValue("kode_store")

	result, err := stock_kasir.Read_Stock_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateHargaStockKasir(c echo.Context) error {
	var Request request_kasir.Update_Stock_Kasir_Request
	var Request_kode request_kasir.Update_Stock_Kasir_Kode_Request

	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

	Request_kode.Kode_barang_kasir = c.FormValue("kode_barang_kasir")

	result, err := stock_kasir.Update_Harga_Stock_Kasir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
