package user_management

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/user_management"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputUserManagement(c echo.Context) error {

	var Request request_kasir.Input_User_Management_Request
	Request.Nama_store = c.FormValue("nama_store")
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := user_management.Input_User_Management(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadUserManagement(c echo.Context) error {
	var Request request_kasir.Read_User_Management_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := user_management.Read_User_Management(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
