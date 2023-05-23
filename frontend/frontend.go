package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"EmailAliasManager/frontend/lib/httpfs"
	"EmailAliasManager/frontend/lib/redirect"
	"EmailAliasManager/frontend/lib/templating"
	"EmailAliasManager/sharedlib/middleware"
)

//go:embed all:templates
var templateEmbed embed.FS

//go:embed all:static
var staticEmbed embed.FS

func Frontend() {
	templateSub, _ := fs.Sub(templateEmbed, "templates")

	// no need to sub because the path is /static
	httpStaticEmbed := http.FS(staticEmbed)

	// register frontend routes
	http.HandleFunc("/login", middleware.UnauthMiddleware(templating.HTMLTemplate("login.html", templateSub)))

	http.HandleFunc("/setup", middleware.UnauthMiddleware(templating.HTMLTemplate("setup.html", templateSub)))

	http.HandleFunc("/home", middleware.AuthMiddleware(templating.HTMLTemplate("home.html", templateSub)))

	http.HandleFunc("/", redirect.Redirect("/home"))

	// serve static css, js and assets
	http.Handle("/static/", http.FileServer(httpfs.FileSystem{Fs: httpStaticEmbed}))
}
