package actions

import (
	"github.com/gobuffalo/buffalo"
)

// EntryHandler is the default function to serve up the entry page.
func EntryHandler(c buffalo.Context) error {
	// dont show the nav bar
	c.Set("no-nav", true)

	return c.Render(200, r.HTML("entry.html"))
}

// HomeHandler is the default function to serve up a home page.
func HomeHandler(c buffalo.Context) error {

	return c.Render(200, r.HTML("home.html"))
}

// CodeHandler is the code for (/code)
func CodeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}

// SoundsHandler stands for (/sounds)
func SoundsHandler(c buffalo.Context) error {
	c.Set("heading", "Sounds")
	return c.Render(200, r.HTML("sounds/sounds.html"))
}

// ResearchHandler is the code for (/research)
func ResearchHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("research/research.html"))
}

// DesignHandler is the code for (/design)
func DesignHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("design/index.html"))
}

// WorkHandler represents (/work)
func WorkHandler(c buffalo.Context) error {
	// Set Name
	c.Set("name", "Obi Knows")
	c.Set("jobs", []string{"Producer", "Designer", "Developer", "Artist", "Researcher"})

	return c.Render(200, r.HTML("work/index.html"))
}

// HealthHandler is the code for (/health)
func HealthHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("health/health.html"))
}

// ContactHandler is the code for (/contact)
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact/index.html"))
}

// CryptoHandler is the code for (/crypto)
func CryptoHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("crypto/index.html"))
}

// FilmsHandler default implementation.
func FilmsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("films/index.html"))
}

// KoboHandler default implementation.
func KoboHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("kobo/index.html"))
}

// SitesHandler default implementation.
func SitesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("sites/index.html"))
}

// AboutHandler default implementation.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}

// DataIndex default implementation.
func DataIndex(c buffalo.Context) error {

	return c.Render(200, r.HTML("data/index.html"))
}
