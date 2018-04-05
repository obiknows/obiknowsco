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
	// Set Name
	c.Set("name", "Obinna Nnodim")
	c.Set("jobs", []string{"Producer", "Designer", "Developer", "Artist", "Researcher"})

	return c.Render(200, r.HTML("work/workpage.html"))
}

// BeatsHandler is a function to serve up
// the Obi Knows portfolio single page app.
func BeatsHandler(c buffalo.Context) error {
	c.Set("heading", "Obi Knows Beats")
	return c.Render(200, r.HTML("beats/beats.html"))
}

// ResearchHandler is a function to serve up
// the Obi Knows Research Blog
func ResearchHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("research/research.html"))
}
