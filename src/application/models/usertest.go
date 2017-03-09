package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type Question struct {
	Question string
	Type string
}


type UserTest struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Title string
	Latitude float32
	Longitude float32
	Questions[] Question
}

func (* UserTest) FindAll() []UserTest {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("usertests")
	var results []UserTest
	err = c.Find(nil).All(&results)
	if err != nil {
		return nil
	} else {
		return results
	}

}

func (ut UserTest) Save() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("usertests")

	err = c.Insert(ut)
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}