package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// WorkHandler is a handler to serve up
// the Obi Knows past work page.
func WorkHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("work/index.html"))
}
