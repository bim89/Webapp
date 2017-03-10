package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

type User struct {
	Name  string
	Email string
}

func (u User) Save(user User) {
	fmt.Println("Name:", user.Name)
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("CEApp").C("user")
	index := mgo.Index{
		Key: []string{"email"},
		Unique: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Insert(&User{u.Name, u.Email})
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}
