package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obiknows/obiknowsco/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
