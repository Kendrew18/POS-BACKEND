package satuan_barang

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/satuan_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSatuanBarang(c echo.Context) error {

	var Request request.Input_Satuan_Barang_Request
	Request.Kode_satuan = c.FormValue("Kode_jenis_barang")
	Request.Nama_satuan = c.FormValue("nama_barang")
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
