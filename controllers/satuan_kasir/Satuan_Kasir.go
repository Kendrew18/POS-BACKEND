package satuan_kasir

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/satuan_kasir"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSatuanKasir(c echo.Context) error {

	var Request request_kasir.Input_Satuan_Kasir_Request
	Request.Nama_satuan = c.FormValue("nama_satuan")

	result, err := satuan_kasir.Input_Satuan_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSatuanKasir(c echo.Context) error {

	result, err := satuan_kasir.Read_Satuan_Barang()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteSatuanKasir(c echo.Context) error {

	var Request request_kasir.Delete_Satuan_Kasir_Request
	Request.Kode_satuan = c.FormValue("kode_satuan")

	result, err := satuan_kasir.Delete_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
