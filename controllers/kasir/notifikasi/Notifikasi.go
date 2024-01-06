package notifikasi

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/notifikasi"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateJumlahMinimal(c echo.Context) error {

	var Request request_kasir.Update_Jumlah_Minimal_Request
	Request.Kode_barang_kasir = c.FormValue("kode_barang_kasir")
	Request.Jumlah_minimal, _ = strconv.ParseFloat(c.FormValue("jumlah_minimal"), 64)

	result, err := notifikasi.Update_Jumlah_Minimal(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadNotifikasi(c echo.Context) error {

	var Request request_kasir.Read_Notifikasi_Kasir_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := notifikasi.Read_Notifikasi(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
