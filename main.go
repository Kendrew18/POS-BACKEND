package main

import (
	"POS-BACKEND/db"
	"POS-BACKEND/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
