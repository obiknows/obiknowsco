package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/obiknows/obiknowsco/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// T is for the translator
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_obiknowsco_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		// Obi Knows Co - Main Site
		app.GET("/", EntryHandler)
		app.GET("/home", HomeHandler)
		app.GET("/code", CodeHandler)
		app.GET("/sounds", SoundsHandler)
		app.GET("/research", ResearchHandler)
		app.GET("/design", DesignHandler)
		app.GET("/work", WorkHandler)
		app.GET("/crypto", CryptoHandler)
		app.GET("/contact", ContactHandler)
		app.GET("/health", HealthHandler)
		app.GET("/films", FilmsHandler)
		app.GET("/kobo", KoboHandler)
		app.GET("/sites", SitesHandler)
		app.GET("/about", AboutHandler)
		// Raw Data & Visualizations
		app.GET("/data", DataIndex)
		// Webstore
		app.GET("/store", StoreIndex)

		// Authentication
		// auth := app.Group("/auth")
		// auth.GET("/", AuthPage)
		// auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		// auth.GET("/{provider}/callback", AuthCallback)

		// Admin Stuffs
		app.GET("/play", PlaygroundHandler)
		app.ServeFiles("/assets", assetsBox)
		app.GET("/routes", RoutesRoutes)

	}

	return app
}
