package actions

import "github.com/gobuffalo/buffalo"

// PlaygroundHandler is the code for the play go round
func PlaygroundHandler(c buffalo.Context) error {
	r.JavaScript("playground/play.js")
	return c.Render(200, r.HTML("playground/play.html"))
}
