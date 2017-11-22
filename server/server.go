package server

import (
	"../config"
	"../db"
)

// Init creates and run NewRouter
func Init() {
	c := config.GetConfig()
	db.InitSession(c.GetString("db.host"))

	defer db.Close()

	r := NewRouter()
	r.Run(c.GetString("app.port"))
}
