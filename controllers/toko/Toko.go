package toko

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/toko"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputToko(c echo.Context) error {
	var Request request.Input_Toko_Request

	Request.Alamat = c.FormValue("alamat")
	Request.Nama_toko = c.FormValue("nama_toko")
	Request.Nomor_telpon = c.FormValue("nomor_telpon")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := toko.Input_Toko(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadToko(c echo.Context) error {
	var Request request.Read_Toko_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := toko.Read_Toko(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func DeleteToko(c echo.Context) error {

	var Request request.Delete_Toko_Request

	Request.Kode_toko = c.FormValue("kode_toko")

	result, err := toko.Delete_Toko(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DropdownNamaToko(c echo.Context) error {

	var Request request.Read_Toko_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := toko.Dropdown_Nama_Toko(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
