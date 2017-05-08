package webapp

import (
	"application/models"
	"testing"
	. "gopkg.in/check.v1"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"io"
)


func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite)TestCheckEmail(c *C) {

	var u = models.User{}

	u.Name = "Bjørn-Inge Morstad"
	u.Email = "bimorstad@gmail.com"
	u.Password = "1234"
	u.Username = "bim"
	u.Gender = "male"
	u.Age = 28
	u.Admin = false
	u.Save()

	_, hasMail := u.CheckEmail("bimorstad@msn.com")
	c.Assert(hasMail, Equals, false)
	u, hasMail = u.CheckEmail("bimorstad@gmail.com")
	c.Assert(hasMail, Equals, true)
	c.Assert(u.Email, Equals, "bimorstad@gmail.com")
}

func (s *MySuite)TestLogin(c *C) {

	body,_ := loginRequest(c, "bimorstad@gmail.com", "1234")
	c.Assert(string(body), Equals, "You have been logged in")
	body2,_ := loginRequest(c, "bimorstad@gmail.com", "12345")
	c.Assert(string(body2), Equals, "Wrong password")
	body3,_ := loginRequest(c, "bimorstad@msn.com", "1234")
	c.Assert(string(body3), Equals, "Email not registered")

}


func (s *MySuite)TestRegister(c *C) {
	body := registerRequest(c, "bimorstad@outlook.com", "bim89", 28, "Male", "1234", "1234")
	u := models.User{}
	json.NewDecoder(body).Decode(&u)
	c.Assert(len(u.UUID), Not(Equals), 0)
	c.Assert(u.Email, Equals, "bimorstad@outlook.com")
}


func loginRequest(c *C, email string, password string) ([]byte, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8001/user/login?email=%s&password=%s", email, password)
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.Fail()
	} else {
		c.Succeed()
	}

	req, _ := client.Do(r)
	return ioutil.ReadAll(req.Body)
}

func registerRequest(c *C, email string, username string, age int, gender string, password string, confirmPassword string) io.ReadCloser {
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8001/user/create?email=%s&username=%s&age=%d&gender=%s&password=%s&confirmPassword=%s",
		email, username, age, gender, password, confirmPassword)
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.Fail()
	} else {
		c.Succeed()
	}

	req, _ := client.Do(r)
	return req.Body
}