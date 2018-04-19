package actions

import "github.com/gobuffalo/buffalo"

// StoreIndex default implementation.
func StoreIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("store/index.html"))
}
