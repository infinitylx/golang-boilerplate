package db

import (
	"sync"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
//var instance *singleton
var once sync.Once

//func GetInstance() *singleton {
//   once.Do(func() {
//        instance = &singleton{}
//    })
//    return instance
//}

// InitSession open mongo db session
func InitSession(host string){
	once.Do(func() {
        s, err := mgo.Dial(host)
        err != nil {
            panic(err)
        }
        s.SetMode(mgo.Monotonic, true)
        session = s
    })
}

// Clone makes and return new session
func Clone() *mgo.Session {
	return session.Clone()
}

// Close session to mongodb.
func Close() {
	session.Close()
}
