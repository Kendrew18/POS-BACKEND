package user

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var user_req request.User_Request
	user_req.Username = c.FormValue("username")
	user_req.Password = c.FormValue("password")

	result, err := user.Login(user_req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func Change_Fifo_Lifo(c echo.Context) error {

	var status_req request.Status_Fifo_Lifo_Request
	status_req.Status, _ = strconv.Atoi(c.FormValue("status"))
	Kode_gudang := c.FormValue("kode_gudang")

	result, err := user.Change_Fifo_Lifo(status_req, Kode_gudang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
