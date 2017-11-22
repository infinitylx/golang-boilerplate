package db

import (
	"gopkg.in/mgo.v2"
)

// Connection DB session store
type Connection struct {
	session *mgo.Session
}

// var Session *mgo.Session
// InitSession open mongo db session
func (conn *Connection) InitSession(host string){
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	conn.session = s
}

// Clone makes and return new session
func (conn *Connection) Clone() *mgo.Session {
	return conn.session.Clone()
}

// Close session to mongodb.
func (conn *Connection) Close() {
	conn.session.Close()
}
