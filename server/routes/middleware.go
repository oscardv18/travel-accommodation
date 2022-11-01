package routes

import (
	"net/http"
)

// middleware for root route
// An middleware need of next function, this the value of return, when info passing of filters in the middleware returning this function and serve of route
func rootMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" && r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "404, not found", http.StatusNotFound)
			w.WriteHeader(http.StatusNotFound)
			// w.Write([]byte("404 Not Found"))
			http.Error(w, "Method is not supported :(.", http.StatusNotFound)
		}
	})
}
