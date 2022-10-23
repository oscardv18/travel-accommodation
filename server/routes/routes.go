package routes

import (
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// home page
	mux.HandleFunc("/", RootHandler)

	return mux
}

// root route handler
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("WELCOME"))
}
