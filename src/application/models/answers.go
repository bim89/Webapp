package models

import (
	"gopkg.in/mgo.v2/bson"
)


type Answer struct {
	Index	int	`json:"index"`
	Score 	int	`json:"score"`
}

type Feedback struct {
	Id    		bson.ObjectId 		`bson:"_id,omitempty"`
	UsertestId 	bson.ObjectId 		`json:"usertestId"`
	Answers[] 	Answer			`json:"answers"`
}
