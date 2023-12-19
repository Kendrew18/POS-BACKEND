package kartu_stock

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/gudang/kartu_stock"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadJenisBarang(c echo.Context) error {

	var Request request.Read_Kartu_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal_1 = c.FormValue("tanggal_1")
	Request.Tanggal_2 = c.FormValue("tanggal_2")
	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Kode_stock = c.FormValue("kode_stock")

	result, err := kartu_stock.Read_Kartu_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
