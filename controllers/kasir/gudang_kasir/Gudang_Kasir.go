package gudang_kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/gudang_kasir"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputGudangKasir(c echo.Context) error {

	var Request request_kasir.Input_Gudang_Kasir_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Alamat = c.FormValue("alamat")
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := gudang_kasir.Input_Gudang_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadGudangKasir(c echo.Context) error {

	var Request request_kasir.Read_Gudang_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := gudang_kasir.Read_Gudang_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func DropdownGudang(c echo.Context) error {

	result, err := gudang_kasir.Dropdown_Gudang()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
