package stock

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/stock"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputStock(c echo.Context) error {

	var Request request.Input_Barang_Request
	Request.Kode_jenis_barang = c.FormValue("Kode_jenis_barang")
	Request.Nama_Barang = c.FormValue("nama_barang")
	Request.Harga_jual, _ = strconv.ParseInt(c.FormValue("harga_jual"), 10, 64)
	Request.Satuan = c.FormValue("satuan")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Input_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadBarang(c echo.Context) error {
	var Request request.Read_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Read_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadStock(c echo.Context) error {

	var Request request.Read_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Read_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
