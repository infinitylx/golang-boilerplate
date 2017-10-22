package db

import (
	"gopkg.in/mgo.v2"
)

// Connection DB session store
type Connection struct {
	session *mgo.Session
}

// Session open mongo db session
// var Session *mgo.Session

// NewConnection creates and saves new session, returns connection to db.
func NewConnection(host string) (conn *Connection) {
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	conn = &Connection{s}

	// Session = s
	return conn
}

// Use make something...
func (conn *Connection) Use(dbName, tableName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)
}

// Clone makes and return session clone
func (conn *Connection) Clone() *mgo.Session {
	return conn.session.Clone()
}

// Close session to mongodb.
func (conn *Connection) Close() {
	conn.session.Close()
	return
}
