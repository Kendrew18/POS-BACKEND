package request

type User_Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Status_Fifo_Lifo_Request struct {
	Status int `json:"status"`
}
