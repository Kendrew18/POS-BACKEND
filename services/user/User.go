package user

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
)

func Login(username string, password string) (response.Response, error) {

	var res response.Response
	var us response.User
	con := db.CreateConGorm().Table("user")

	err := con.Select("id_user", "status").Where("username =? and password =?", username, password).Scan(&us).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		us.Id_user = ""
		res.Data = us

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = us
	}

	fmt.Println()

	return res, nil

}
