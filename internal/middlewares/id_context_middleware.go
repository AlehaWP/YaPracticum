package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Key string

// UrlCtx for parameter transfer without direct access to router
func URLCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), Key("id"), chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
