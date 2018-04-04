package actions

import "github.com/gobuffalo/buffalo"

// RoutesRoutes default implementation.
func RoutesRoutes(c buffalo.Context) error {
	return c.Render(200, r.HTML("routes/routes.html"))
}
