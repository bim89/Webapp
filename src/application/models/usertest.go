package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	// "errors"
)

type Question struct {
	Question string 	`json:"question"`
	Type string 		`json:"type"`
}


type UserTest struct {
	Id    bson.ObjectId 	`bson:"_id,omitempty"`
	Title string 		`json:"title"`
	Latitude float32 	`json:"latitude"`
	Longitude float32 	`json:"longitude"`
	Questions[] Question 	`json:"questions"`
}

func (* UserTest) FindId(id string) UserTest {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("usertests")
	ut := UserTest{}
	if bson.IsObjectIdHex(id) {

		err = c.FindId(bson.ObjectIdHex(id)).One(&ut)
		if err != nil {
			// return ut, err
		}

		// return ut, nil

	} else {
		// return ut, errors.New("Not a valid Object ID")
	}

	return ut


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

func (* UserTest) Delete(id string) (string, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("usertests")

	if bson.IsObjectIdHex(id) {
		err = c.RemoveId(bson.ObjectIdHex(id))

		if err != nil {
			return err.Error(), err
		} else {
			return "User Test was deleted", err
		}
	} else {
		return "Not a valid input", err
	}

}