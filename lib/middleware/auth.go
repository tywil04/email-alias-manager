package middleware

import (
	"context"
	"net/http"

	"EmailAliasManager/lib/cryptography"
	"EmailAliasManager/lib/db"
)

func AuthMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return AttachContext(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		users := []db.User{}
		db.DB.Preload("Aliases").Preload("Domains").Preload("Aliases.Domain").Find(&users)

		if len(users) == 0 {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		if !cryptography.CompareTokens(token.Value, users[0].StrengthendToken, users[0].TokenSalt) {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "authedUser", users[0]))

		next(w, r)
	})
}

func UnauthMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return AttachContext(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			next(w, r)
			return
		}

		users := []db.User{}
		db.DB.Find(&users)

		if len(users) == 0 {
			next(w, r)
			return
		}

		if !cryptography.CompareTokens(token.Value, users[0].StrengthendToken, users[0].TokenSalt) {
			next(w, r)
			return
		}

		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	})
}
