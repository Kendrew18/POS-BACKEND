package jenis_barang

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/jenis_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputJenisBarang(c echo.Context) error {

	var Request request.Input_Jenis_Barang_Request
	Request.Nama_Jenis_Barang = c.FormValue("nama_jenis_barang")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := jenis_barang.Input_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadJenisBarang(c echo.Context) error {

	var Request request.Read_Jenis_Barang_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := jenis_barang.Read_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func DeleteJenisBarang(c echo.Context) error {
	var Request request.Delete_Jenis_Barang_Request
	Request.Kode_jenis_barang = c.FormValue("kode_jenis_barang")

	result, err := jenis_barang.Delete_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
