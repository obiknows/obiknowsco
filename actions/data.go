package actions

import "github.com/gobuffalo/buffalo"

// DataIndex default implementation.
func DataIndex(c buffalo.Context) error {

	return c.Render(200, r.HTML("data/index.html"))
}
