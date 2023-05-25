package middleware

import (
	"context"
	"net/http"
	"time"
)

func AttachContext(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		already := r.Context().Value("context")
		if already == nil || !already.(bool) {
			ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
			defer cancel()
			r = r.WithContext(context.WithValue(ctx, "context", true))
		}
		next(w, r)
	}
}
