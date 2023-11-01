package user

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var Request request.User_Request
	Request.Username = c.FormValue("username")
	Request.Password = c.FormValue("password")

	result, err := user.Login(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func Change_Fifo_Lifo(c echo.Context) error {

	var Request request.Status_Fifo_Lifo_Request
	Request.Status, _ = strconv.Atoi(c.FormValue("status"))
	Kode_gudang := c.FormValue("kode_gudang")

	result, err := user.Change_Fifo_Lifo(Request, Kode_gudang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
