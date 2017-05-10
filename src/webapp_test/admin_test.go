package webapp

import (
	"application/models"
	. "gopkg.in/check.v1"
	"github.com/satori/go.uuid"
)

type AdminSuite struct{}

var _ = Suite(&AdminSuite{})


func (adm *AdminSuite) TestSave(c *C) {
	a := models.Admin{}
	a.Email = "bimorstad@gmail.com"
	a.Password = "1234"
	a.Store = "Joker Nord"
	a.UUID = uuid.NewV4().String()
	a.Logo_image = "some_url.jpg"
	a.Save()


	c.Assert(len(a.Id), Not(Equals), 0)
}