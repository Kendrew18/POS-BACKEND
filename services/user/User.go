package user

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
)

func Login(user request.User_Request) (response.Response, error) {

	var res response.Response
	var us response.User_Response
	con := db.CreateConGorm().Table("user")

	err := con.Select("id_user", "status", "kode_gudang").Where("username =? and password =?", user.Username, user.Password).Scan(&us).Error

	fmt.Println(err)

	if err != nil || us.Id_user == "" {
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
