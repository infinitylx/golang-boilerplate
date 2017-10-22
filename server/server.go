package server

import (
	"../config"
	"../db"
)

// Init creates and run NewRouter
func Init() {
	c := config.GetConfig()

	var dbConnect = db.NewConnection(c.GetString("db.host"))

	defer dbConnect.Close()

	r := NewRouter(*dbConnect)
	r.Run(c.GetString("app.port"))
}
