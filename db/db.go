package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	C       *mgo.Collection
)

func Init() {
	session, err := mgo.Dial("localhost")
	Session = session
	if err != nil {
		panic(err)
	}
	err = Session.Ping()
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	C = Session.DB("test").C("stocks")
	// defer session.Close()
}

func AddData(s ...interface{}) {
	err := C.Insert(s...)
	if err != nil {
		log.Fatal(err)
	}
}
func CloseDB() {
	Session.Close()
}
