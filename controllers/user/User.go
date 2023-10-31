package user

import (
	"POS-BACKEND/models/request"
	"POS-BACKEND/services/user"
	"net/http"

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
