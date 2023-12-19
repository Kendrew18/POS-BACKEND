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

func Change_Fifo_Lifo(status request.Status_Fifo_Lifo_Request, kode_gudang string) (response.Response, error) {
	var res response.Response

	//0=fifo
	//1=lifo

	con := db.CreateConGorm().Table("user")

	err := con.Where("kode_gudang =?", kode_gudang).Update("status", status.Status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = status
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}
