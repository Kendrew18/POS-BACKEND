package stock_keluar

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/stock_keluar"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Input Stock Keluar
func InputStockKeluar(c echo.Context) error {

	var Request request.Input_Stock_Keluar_Request
	var Request_barang request.Input_Barang_Stock_Keluar_Request

	Request.Tanggal_keluar = c.FormValue("tanggal_stock_keluar")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode_toko = c.FormValue("kode_toko")
	Request.Nama_penanggung_jawab = c.FormValue("penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_jual = c.FormValue("harga_jual")

	result, err := stock_keluar.Input_Stock_Keluar(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

//Read Stock Masuk
func ReadStockKeluar(c echo.Context) error {
	var Request request.Read_Stock_Keluar_Request
	var Request_filter request.Filter_Stock_Keluar_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_toko = c.FormValue("kode_toko")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

	result, err := stock_keluar.Read_Stock_Keluar(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
