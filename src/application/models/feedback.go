package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)


type Answer struct {
	Index	int	`json:"index"`
	Score 	int	`json:"score"`
}

type Feedback struct {
	Id    		bson.ObjectId 		`bson:"_id,omitempty"`
	UsertestId 	bson.ObjectId 		`bson:",omitempty"`
	Answers[] 	Answer			`json:"answers"`
}


func (f Feedback) Save() (string, error) {

	c, session := getCollection("feedback")
	defer session.Close()

	err := c.Insert(f)
	if mgo.IsDup(err) {
		return "Duplicate Insert", err
	} else {
		return "Your feedback was added", err
	}
}

