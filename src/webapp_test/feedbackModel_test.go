
package webapp_test

import (
	"testing"
	. "gopkg.in/check.v1"
	"application/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)


func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite)TestSave(c *C) {
	a1 := models.Answer{}
	a1.Index = 0
	a1.Score = 3

	a2 := models.Answer{}
	a2.Index = 1
	a2.Score = 3

	a3 := models.Answer{}
	a3.Index = 2
	a3.Score = 3
	answers := []models.Answer{a1, a2, a3}

	f := models.Feedback{}
	f.UsertestId = bson.ObjectIdHex("58c2814bef5f386c36ce0589")
	f.Answers = answers
	msg, err := f.Save()
	if err != nil {
		c.Fail()
	} else {
		c.Assert(msg, Equals, "Your feedback was added")
	}
}


func (s *MySuite)TestCreate(c *C) {
	_, err := http.NewRequest("POST", "/feedback/create", nil)
	if err != nil {
		c.Fail()
	} else {
		c.Succeed()
	}
}

