package main

import (
	"AuthTachegando/internal/db"
	"AuthTachegando/routes"
)

func main() {
	db.Init()
	r := routes.SetupRouter()
	r.Run(":8080")
}
