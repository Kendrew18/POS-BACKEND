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

func InputAuditStock(c echo.Context) error {

	var Request request.Input_Audit_stock_Request
	var Request_detail request.Input_Detail_Audit_stock_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal = c.FormValue("tanggal_sekarang")
	Request.Kode_stock = c.FormValue("kode_stock")

	Request_detail.Kode_barang_keluar_masuk = c.FormValue("kode_barang_keluar_masuk")
	Request_detail.Tanggal_masuk = c.FormValue("tanggal_masuk")
	Request_detail.Stock_dalam_sistem = c.FormValue("stock_dalam_sistem")
	Request_detail.Stock_rill = c.FormValue("stock_rill")
	Request_detail.Selisih_stock = c.FormValue("selisih_stock")

	result, err := audit.Input_Audit_Stock(Request, Request_detail)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadAuditStock(c echo.Context) error {

	var Request request.Read_Audit_Stock
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal = c.FormValue("tanggal")

	result, err := audit.Read_Audit_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
