package kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/kasir"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadMenuKasir(c echo.Context) error {

	var Request request_kasir.Read_Stock_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")
	Request.Kode_store = c.FormValue("kode_store")

	result, err := kasir.Read_Menu_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
