package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is t default function to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// WorkHandler is a function to serve up
// the Obi Knows portfolio single page app.
func WorkHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("work/index.html"))
}

// ResearchHandler is a function to serve up
// the Obi Knows Research Blog
func ResearchHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("research/index.html"))
}
