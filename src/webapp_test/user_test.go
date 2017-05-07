package webapp

import (
	"application/models"
	"testing"
	. "gopkg.in/check.v1"
	"net/http"
	"io/ioutil"
	"fmt"
)


func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite)TestCheckEmail(c *C) {

	var u = models.User{}
	/*
	u.Name = "Bjørn-Inge Morstad"
	u.Email = "bimorstad@gmail.com"
	u.Password = "1234"
	u.Username = "bim"
	u.Gender = "male"
	u.Age = 28
	u.Admin = false
	u.Save()
	*/

	_, hasMail := u.CheckEmail("bimorstad@outlook.com")
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
	body3,_ := loginRequest(c, "bimorstad@outlook.com", "1234")
	c.Assert(string(body3), Equals, "Email not registered")

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