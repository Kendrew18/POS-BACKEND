package bentuk_retur

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/bentuk_retur"

	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSatuanKasir(c echo.Context) error {

	var Request request_kasir.Input_Bentuk_Retur_request
	Request.Nama_bentuk_retur = c.FormValue("nama_bentuk_retur")

	result, err := bentuk_retur.Input_Bentuk_Retur(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadBentukRetur(c echo.Context) error {

	result, err := bentuk_retur.Read_Bentuk_Retur()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
