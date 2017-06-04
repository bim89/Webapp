
package webapp_test

import (
	. "gopkg.in/check.v1"
	"application/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"io"
	"bytes"
	"io/ioutil"
	"time"
)




type FeedbackSuite struct{}

var _ = Suite(&FeedbackSuite{})

func createFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"usertestid": "5911e9f635bafe7b705d0491", "answers": [{ "index": 0, "score": 3 }, { "index": 1, "score": 2 }, { "index": 2, "score": 5 }]}`)
}

func (s *FeedbackSuite)TestSave(c *C) {
	a1 := models.Answer{}
	a1.Index = 0
	a1.Score = 4


	answers := []models.Answer{a1}

	f := models.Feedback{}
	f.UsertestId = bson.ObjectIdHex("593269b4122362289b6ee048")
	f.AnsweredBy = "bimorstad@gmail.com"
	f.Answers = answers
	f.Date = makeTimestampMilli("2017-06-09T15:05:41+02:00")

	msg, err := f.Save()
	if err != nil {
		c.Fail()
	} else {
		c.Assert(msg, Equals, "Your feedback was added")
	}

}



func (s *FeedbackSuite)TestCreate(c *C) {

	json :=  `{"usertestid": "5911e9f635bafe7b705d0491", "answers": [{ "index": 0, "score": 2 }, { "index": 1, "score": 4 }, { "index": 2, "score": 2 }]}`

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


func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func makeTimestampMilli(dateVal string) int64 {
	t,_ := time.Parse(time.RFC3339, dateVal)
	return unixMilli(t)
}

