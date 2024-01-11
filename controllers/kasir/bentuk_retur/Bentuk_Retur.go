package bentuk_retur

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/bentuk_retur"

	"net/http"

	"github.com/labstack/echo/v4"
)

func InputBentukRetur(c echo.Context) error {

	var Request request_kasir.Input_Bentuk_Retur_Request
	Request.Nama_bentuk_retur = c.FormValue("nama_bentuk_retur")
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := bentuk_retur.Input_Bentuk_Retur(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadBentukRetur(c echo.Context) error {

	var Request request_kasir.Read_Barang_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := bentuk_retur.Read_Bentuk_Retur(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func DeleteBentukRetur(c echo.Context) error {

	var Request request_kasir.Delete_Bentuk_Retur_Request
	Request.Kode_bentuk_retur = c.FormValue("kode_bentuk_retur")

	result, err := bentuk_retur.Delete_Bentuk_Retur(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
