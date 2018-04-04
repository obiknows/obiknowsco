package actions

import "github.com/gobuffalo/buffalo"

// PlayIndex default implementation.
func PlayIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("play/index.html"))
}
