package kartu_stock

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/kartu_stock"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadJenisBarang(c echo.Context) error {

	var Request request.Read_Kartu_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal = c.FormValue("tanggal")

	result, err := kartu_stock.Read_Kartu_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
