package server

import (
	"../config"
	"../db"
)

// Init creates and run NewRouter
func Init() {
	c := config.GetConfig()
	mmgo = db.Connection

	mmgo.InitSession(c.GetString("db.host"))

	defer mmgo.Close()

	r := NewRouter()
	r.Run(c.GetString("app.port"))
}
