
package webapp_test

import (
	"testing"
	. "gopkg.in/check.v1"
	"application/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"io"
	"bytes"
	"io/ioutil"
)


func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func createFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"usertestid": "58c2814bef5f386c36ce0584", "answers": [{ "index": 0, "score": 3 }, { "index": 1, "score": 2 }, { "index": 2, "score": 5 }]}`)
}

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

	json :=  `{"usertestid": "58c14c5737fb32e6c63da5af", "answers": [{ "index": 0, "score": 2 }, { "index": 1, "score": 4 }, { "index": 2, "score": 2 }]}`

	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://localhost:8001/feedback/create", bytes.NewBufferString(json))
	if err != nil {
		c.Fail()
	} else {
		c.Succeed()
	}

	req, _ := client.Do(r)

	body, err := ioutil.ReadAll(req.Body)
	c.Assert(string(body), Equals, "Your feedback was added")

}

