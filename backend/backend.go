package backend

import (
	"net/http"

	"EmailAliasManager/backend/routes/v1/user/aliases"
	"EmailAliasManager/backend/routes/v1/user/data"
	"EmailAliasManager/backend/routes/v1/user/domains"
	"EmailAliasManager/backend/routes/v1/user/setup"
	"EmailAliasManager/lib/middleware"
)

func Backend() {
	http.HandleFunc("/api/v1/user/setup", setup.Post)

	http.HandleFunc("/api/v1/user/aliases", middleware.MethodRouter(middleware.MethodMap{
		Post:   middleware.AuthMiddleware(aliases.Post),
		Patch:  middleware.AuthMiddleware(aliases.Patch),
		Delete: middleware.AuthMiddleware(aliases.Delete),
	}))

	http.HandleFunc("/api/v1/user/domains", middleware.MethodRouter(middleware.MethodMap{
		Post:   middleware.AuthMiddleware(domains.Post),
		Patch:  middleware.AuthMiddleware(domains.Patch),
		Delete: middleware.AuthMiddleware(domains.Delete),
	}))

	http.HandleFunc("/api/v1/user/data", middleware.MethodRouter(middleware.MethodMap{
		Get:  middleware.AuthMiddleware(data.Get),
		Post: middleware.AuthMiddleware(data.Post),
	}))
}
