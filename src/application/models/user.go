package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type TempUser struct {
	Email 			string		`json:"email"`
	Username		string		`json:"username"`
	Age 			int		`json:"age"`
	Gender			string		`json:"gender"`
	Password 		string		`json:"password"`
	ConfirmPassword 	string		`json:"confirm_password"`
}

type User struct {
	Id    		bson.ObjectId 	`bson:"_id,omitempty"`
	Email 		string		`json:"email"`
	Password 	string		`json:"-"`
	Username	string		`json:"username"`
	Age 		int		`json:"age"`
	Gender		string		`json:"gender"`
	UUID  		string		`json:"-"`
}


func (u User) CheckEmail(email string) (User, bool) {
	c, session := getCollection("user")
	defer session.Close()
	err := c.Find(bson.M{"email": email}).One(&u)

	if err == nil {
		if u.Email == email {
			return u, true
		}
	} else {
		log.Println(err)
	}

	return u, false
}

func (u User) Update(user User) bool {
	c, session := getCollection("user")
	defer session.Close()

	err := c.UpdateId(u.Id, u)
	if err != nil {
		log.Panic(err)
	} else {
		return true
	}

	return false
}

func (u User) Save() {
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

	err = c.Insert(u)
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}
