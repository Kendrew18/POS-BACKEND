package satuan_barang

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/satuan_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSatuanBarang(c echo.Context) error {

	var Request request.Input_Satuan_Barang_Request
	Request.Nama_satuan_barang = c.FormValue("nama_satuan")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := satuan_barang.Input_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSatuanBarang(c echo.Context) error {

	var Request request.Read_Satuan_Barang_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := satuan_barang.Read_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteSatuanBarang(c echo.Context) error {

	var Request request.Delete_Satuan_Barang_Request
	Request.Kode_satuan_barang = c.FormValue("kode_satuan_barang")

	result, err := satuan_barang.Delete_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
