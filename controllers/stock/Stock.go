package stock

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/stock"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputStock(c echo.Context) error {

	var stock_req request.Input_Barang_Request
	stock_req.Kode_jenis_barang = c.FormValue("Kode_jenis_barang")
	stock_req.Nama_Barang = c.FormValue("nama_barang")
	stock_req.Harga_jual, _ = strconv.ParseInt(c.FormValue("harga_jual"), 10, 64)
	stock_req.Satuan = c.FormValue("satuan")
	stock_req.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock.Input_Barang(stock_req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
