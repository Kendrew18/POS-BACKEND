package user_management

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/user_management"
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

func DeleteUserManagement(c echo.Context) error {
	var Request request_kasir.Delete_User_Management_Request
	Request.Kode_store = c.FormValue("kode_store")

	result, err := user_management.Delete_User_Management(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
