package audit

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/audit"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadDataAwalAuditStock(c echo.Context) error {

	var Request request.Read_Data_Awal_Audit_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal_sekarang = c.FormValue("tanggal_sekarang")

	result, err := audit.Read_Data_Awal_Audit_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
