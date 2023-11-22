package stock

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/stock"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputBarang(c echo.Context) error {

	var Request request.Input_Barang_Request
	Request.Kode_jenis_barang = c.FormValue("kode_jenis_barang")
	Request.Nama_Barang = c.FormValue("nama_barang")
	Request.Harga_jual = c.FormValue("harga_jual")
	Request.Kode_satuan_barang = c.FormValue("kode_satuan_barang")
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
	Request.Kode_jenis_barang = c.FormValue("kode_jenis_barang")

	result, err := stock.Read_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadStock(c echo.Context) error {

	var Request request.Read_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Kode_jenis_barang = c.FormValue("kode_jenis_barang")

	result, err := stock.Read_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteBarang(c echo.Context) error {
	var Request request.Delete_Barang_Request
	Request.Kode_stock = c.FormValue("kode_stock")

	result, err := stock.Delete_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func Detailstock(c echo.Context) error {
	var Request request.Read_Detail_Stock
	Request.Kode_stock = c.FormValue("kode_stock")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Detail_stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DropdownStock(c echo.Context) error {
	var Request request.Dropdown_Stock_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Dropdown_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
