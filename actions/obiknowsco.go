package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is t default function to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// WorkHandler represents (/work)
func WorkHandler(c buffalo.Context) error {
	// Set Name
	c.Set("name", "Obi Knows")
	c.Set("jobs", []string{"Producer", "Designer", "Developer", "Artist", "Researcher"})

	return c.Render(200, r.HTML("work/workpage.html"))
}

// BeatsHandler stands for (/beats)
func BeatsHandler(c buffalo.Context) error {
	c.Set("heading", "Obi Knows Beats")
	return c.Render(200, r.HTML("beats/beats.html"))
}

// ResearchHandler is the code for (/research)
func ResearchHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("research/research.html"))
}

// ContactHandler is the code for (/contact)
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact/contactpage.html"))
}

// CodeHandler is the code for (/code)
func CodeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}

// CryptoHandler is the code for (/crypto)
func CryptoHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("crypto/index.html"))
}
