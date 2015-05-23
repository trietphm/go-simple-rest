package resource

import (
	"gopkg.in/mgo.v2"
)

var (
	mongoSession *mgo.Session
	db           string
)

func InitDB() error {
	var err error
	if mongoSession, err = mgo.Dial("127.0.0.1"); err != nil {
		return err
	}
	db = "demo"
	return nil
}

func collection(c string) *mgo.Collection {
	return mongoSession.DB(db).C(c)
}

func closeDB() {
	if mongoSession != nil {
		mongoSession.Close()
	}
}
