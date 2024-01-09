package request_barang_gudang

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/gudang/request_barang_gudang"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ReadRequestBarangKasirStock(c echo.Context) error {
	var Request request.Read_Request_Barang_Kasir_Stock_Request
	var Request_Filter request.Read_Filter_Request_Barang_Kasir_Stock_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_Filter.Tanggal_1 = c.FormValue("tanggal")

	result, err := request_barang_gudang.Read_Request_Barang_Kasir_Stock(Request, Request_Filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusRequestBarangKasir(c echo.Context) error {
	var Request request_kasir.Update_Status_Request_Barang_Kasir
	var Request_kode request.Update_Status_Kode_Request_Barang_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	Request_kode.Kode_request_barang_kasir = c.FormValue("kode_request_barang_kasir")
	Request_kode.Kode_gudang = c.FormValue("kode_gudang")
	Request_kode.Kode_user = c.FormValue("kode_user")
	Request_kode.Kode_nota = c.FormValue("kode_nota")

	result, err := request_barang_gudang.Update_Status_Request_Barang_Kasir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
