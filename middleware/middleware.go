package middleware

import (
	"net/http"

	"github.com/livghit/gophosaurus/views"
)

func NotFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusNotFound)
	views.NotFoundView().Render(r.Context(), w)
}
