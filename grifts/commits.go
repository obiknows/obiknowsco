package grifts

import (
	"fmt"

	. "github.com/markbates/grift/grift"
)

var _ = Namespace("code", func() {

	Desc("commits", "Gets json array of a year of github commit data for the calendar")
	Add("commits", func(c *Context) error {
		fmt.Println("Lets get these commits...")
		return nil
	})

})
