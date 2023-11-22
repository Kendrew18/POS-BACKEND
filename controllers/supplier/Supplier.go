package supplier

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/supplier"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Input Supplier
func InputSupplier(c echo.Context) error {
	var Request request.Input_Supplier_Request
	var Request_barang request.Input_Barang_Supplier_Request

	Request.Nama_supplier = c.FormValue("nama_supplier")
	Request.Nomor_telpon = c.FormValue("nomor_telpon")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")

	result, err := supplier.Input_Supplier(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

//Read Supplier
func ReadSupplier(c echo.Context) error {
	var Request request.Read_Supplier_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := supplier.Read_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

//Dropdown Nama Supplier
func DropdownNamaSupplier(c echo.Context) error {
	var Request request.Read_Supplier_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := supplier.Dropdown_Nama_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteSupplier(c echo.Context) error {

	var Request request.Delete_Supplier_Request

	Request.Kode_supplier = c.FormValue("kode_supplier")

	result, err := supplier.Delete_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DropdownBarangSupplier(c echo.Context) error {
	var Request request.Read_Barang_Supplier_Request

	Request.Kode_supplier = c.FormValue("kode_supplier")

	result, err := supplier.Dropdown_Barang_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
