package backend

import (
	"net/http"

	"EmailAliasManager/backend/routes/v1/user/aliases"
	"EmailAliasManager/backend/routes/v1/user/domains"
	"EmailAliasManager/backend/routes/v1/user/setup"
	"EmailAliasManager/sharedlib/middleware"
)

func Backend() {
	http.HandleFunc("/api/v1/user/setup", setup.Post)

	http.HandleFunc("/api/v1/user/aliases", middleware.MethodRouter(middleware.MethodMap{
		Post:  middleware.AuthMiddleware(aliases.Post),
		Patch: middleware.AuthMiddleware(aliases.Patch),
	}))

	http.HandleFunc("/api/v1/user/domains", middleware.MethodRouter(middleware.MethodMap{
		Post:  middleware.AuthMiddleware(domains.Post),
		Patch: middleware.AuthMiddleware(domains.Patch),
	}))
}
