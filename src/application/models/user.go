package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

type User struct {
	Name  string
	Email string
	Password string
	UUID  string
}

func (u User) Save(user User) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()



	c := session.DB("CEApp").C("user")
	index := mgo.Index{
		Key: []string{"email"},
		Unique: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Insert(&User{u.Name, u.Email, u.Password, u.UUID})
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}
