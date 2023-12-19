package pre_order

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/gudang/pre_order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputPreOrder(c echo.Context) error {

	var Request request.Input_Pre_Order_Request
	var Request_barang request.Input_Barang_Pre_Order_Request

	Request.Tanggal = c.FormValue("tanggal_pre_order")
	Request.Kode_nota = c.FormValue("kode_nota")
	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Nama_penanggung_jawab = c.FormValue("nama_penanggung_jawab")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Tanggal_kadalurasa = c.FormValue("tanggal_kadalurasa")
	Request_barang.Jumlah_barang = c.FormValue("jumlah_barang")
	Request_barang.Harga_pokok = c.FormValue("harga_pokok")

	result, err := pre_order.Input_Pre_Order(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadPreOrder(c echo.Context) error {

	var Request request.Read_Pre_Order_Request
	var Request_filter request.Read_Pre_Order_Filter_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request_filter.Kode_supplier = c.FormValue("kode_supplier")
	Request_filter.Tanggal_1 = c.FormValue("tanggal_1")
	Request_filter.Tanggal_2 = c.FormValue("tanggal_2")

	result, err := pre_order.Read_Pre_Order(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdatePreOrder(c echo.Context) error {

	var Request_kode request.Update_Pre_Order_Kode_Request
	var Request request.Update_Pre_order_Request

	Request_kode.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")
	Request.Tanggal_kadaluarsa = c.FormValue("tanggal_kadaluarsa")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

	result, err := pre_order.Update_Pre_Order(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeletePreOrder(c echo.Context) error {

	var Request_kode request.Update_Pre_Order_Kode_Request

	Request_kode.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")

	result, err := pre_order.Delete_Pre_Order(Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusPreOrder(c echo.Context) error {

	var Request_kode request.Kode_Pre_Order_Request
	var Request request.Update_Status_Pre_Order_Request

	Request_kode.Kode_pre_order = c.FormValue("kode_pre_order")
	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	result, err := pre_order.Update_Status_Pre_Order(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
