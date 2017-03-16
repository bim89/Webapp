package models

import (
	"gopkg.in/mgo.v2"
)


func getCollection(colllection string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session.DB("CEApp").C(colllection), session

}
