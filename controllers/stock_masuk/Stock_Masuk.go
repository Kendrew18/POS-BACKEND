package stock_masuk

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/stock_masuk"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//Input Stock Masuk
func InputStockMasuk(c echo.Context) error {

	var Request request.Input_Stock_Masuk_Request
	var Request_barang request.Input_Barang_Stock_Masuk_Request

	Request.Tanggal = c.FormValue("tanggal_stock_masuk")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode = c.FormValue("kode_supplier")
	Request.Nama_penanggung_jawab = c.FormValue("nama_penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Tanggal_kadalurasa = c.FormValue("tanggal_kadalurasa")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_pokok = c.FormValue("harga_pokok")

	result, err := stock_masuk.Input_Stock_Masuk(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

//Read Stock Masuk
func ReadStockMasuk(c echo.Context) error {
	var Request request.Read_Stock_Masuk_Request
	var Request_filter request.Read_Stock_Masuk_Filter_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_supplier = c.FormValue("kode_supplier")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

	result, err := stock_masuk.Read_Stock_Masuk(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateBarangStockMasuk(c echo.Context) error {
	var Request request.Update_Stock_Masuk_Request
	var Request_kode request.Update_Stock_Masuk_Kode_Request

	Request_kode.Kode_barang_keluar_masuk = c.FormValue("kode_barang_keluar_masuk")
	Request.Tanggal_kadaluarsa = c.FormValue("tanggal_kadaluarsa")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

	result, err := stock_masuk.Update_Barang_Stock_Masuk(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteBarangStockMasuk(c echo.Context) error {
	var Request_kode request.Update_Stock_Masuk_Kode_Request

	Request_kode.Kode_barang_keluar_masuk = c.FormValue("kode_barang_keluar_masuk")

	result, err := stock_masuk.Delete_Barang_Stock_Masuk(Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
