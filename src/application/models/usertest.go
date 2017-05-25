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
	Choices[] Choices	`json:"choices"`
}

type Choices struct {
	Answer string 		`json:"answer"`
}


type UserTest struct {
	Id    		bson.ObjectId 	`bson:"_id,omitempty"`
	Title 		string 		`json:"title"`
	Email	 	string		`json:"email"`
	Latitude 	float32 	`json:"latitude"`
	Longitude 	float32 	`json:"longitude"`
	Questions[] 	Question 	`json:"questions"`
	Feedback[] 	Feedback 	`json:"feedback"`
	Admin		Admin		`json:"admin"`
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
		} else {
			feedback := session.DB("CEApp").C("feedback")
			results := []Feedback{}
			err = feedback.Find(bson.M{"usertestid": bson.ObjectIdHex(id)}).All(&results)
			if err != nil {
				log.Panic(err)
			}
			log.Println(len(results))
			ut.Feedback = results
			for index, elem := range ut.Feedback {
				user, hasmail := User{}.CheckEmail(elem.AnsweredBy)
				if (hasmail) {
					ut.Feedback[index].User = user;
				}
			}
		}
	} else {
		// return ut, errors.New("Not a valid Object ID")
	}

	return ut
}

func (* UserTest) FindAll(email string, withFeedback bool) []UserTest {
	c, session := getCollection("usertests")
	defer session.Close()

	var results []UserTest

	if (email != "") {
		c.Find(bson.M{"email": email}).All(&results)
		// admin := Admin{}.GetByEmail(results[0].Email)
		// results[0].Admin = admin
	} else {
		c.Find(nil).All(&results)
		// admin := Admin{}.GetByEmail(results[0].Email)
		// results[0].Admin = admin
	}

	if withFeedback {
		feedColl, session := getCollection("feedback")
		defer session.Close()
		for i := 0; i < len(results); i++ {
			var feedback []Feedback
			if err := feedColl.Find(bson.M{"usertestid": results[i].Id }).All(&feedback); err != nil {
				log.Panic(err.Error())
			} else {
				println(feedback)
				if (feedback != nil) {
					results[i].Feedback = feedback
					for index, elem := range feedback {
						user, hasmail := User{}.CheckEmail(elem.AnsweredBy)
						if (hasmail) {
							results[i].Feedback[index].User = user;
						}
					}
				}
			}
		}
	}


	return results

}

func (ut UserTest) Save() {

	c, session := getCollection("usertests")
	defer session.Close()

	err := c.Insert(ut)
	if mgo.IsDup(err) {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
}

func (* UserTest) Delete(id string) (string, error) {

	c, session := getCollection("usertests")
	defer session.Close()

	if bson.IsObjectIdHex(id) {
		err := c.RemoveId(bson.ObjectIdHex(id))

		if err != nil {
			return err.Error(), err
		} else {
			return "User Test was deleted", err
		}
	} else {
		return "Not a valid input", nil
	}

}