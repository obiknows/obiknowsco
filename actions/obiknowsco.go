package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/polera/publicip"
)

// HomeHandler is t default function to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	myIPAddr, _ := publicip.GetIP()

	c.Set("my-ip", myIPAddr)
	return c.Render(200, r.HTML("index.html"))
}

// WorkHandler represents (/work)
func WorkHandler(c buffalo.Context) error {
	// Set Name
	c.Set("name", "Obi Knows")
	c.Set("jobs", []string{"Producer", "Designer", "Developer", "Artist", "Researcher"})

	return c.Render(200, r.HTML("work/index.html"))
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

// HealthHandler is the code for (/health)
func HealthHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("health/health.html"))
}

// ContactHandler is the code for (/contact)
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact/index.html"))
}

// CodeHandler is the code for (/code)
func CodeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("code/index.html"))
}

// CryptoHandler is the code for (/crypto)
func CryptoHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("crypto/index.html"))
}

// SebiHandler default implementation.
func SebiHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("sebi/index.html"))
}

// FilmsHandler default implementation.
func FilmsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("films/index.html"))
}

// KoboHandler default implementation.
func KoboHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("kobo/index.html"))
}

// AboutHandler default implementation.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}
