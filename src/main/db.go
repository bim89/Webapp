package main
import (

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type Person struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}


func dbHandler() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("persons")

	err = c.Insert(
		bson.M{"name": "Ale", "phone": "+55 53 8116 9639"},
		bson.M{"name": "Pete", "phone": "+55 53 8556 9544"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Id.Hex())
}

