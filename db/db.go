package db

import (
	"gopkg.in/mgo.v2"
	"sync"
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
func InitSession(host string) {
	once.Do(func() {
		s, err := mgo.Dial(host)
		if err != nil {
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

// Get runs given query and return json
func Get(query string, col string, result interface{}) err error {
	s := Clone()
	defer s.Close()

	var result []User
	err := s.DB("devdb").C(col).Find(nil).All(&result)

    // use named return
    // return err
    return
}

