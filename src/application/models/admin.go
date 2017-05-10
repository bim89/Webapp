package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

type Admin struct {
	Id    		bson.ObjectId 	`bson:"_id,omitempty"`
	Email 		string		`json:"email"`
	Password 	string		`json:"-"`
	Store		string 		`json:"store"`
	Logo_image	string		`json:"logo_image"`
	UUID		string		`json:"-"`
}



func (a Admin) GetByEmail(email string) Admin {
	c, session := getCollection("admin")
	defer session.Close()

	err := c.Find(bson.M{"email": email}).One(&a)
	if err == nil {
		if a.Email == email {
			return a
		}
	} else {
		log.Println(err)
	}

	return a
}

func (a *Admin) Save() {
	c, session := getCollection("admin")
	defer session.Close()

	index := mgo.Index{
		Key: []string{"email"},
		Unique: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		log.Panic(err)
	}

	err = c.Insert(a)
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}
