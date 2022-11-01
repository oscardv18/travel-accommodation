package routes

import (
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// home page
	routeRoot := http.HandlerFunc(rootHandler)
	mux.Handle("/", rootMiddleware(routeRoot))

	return mux
}

// root route handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hello from home path"))
	// http.ServeFile(w, r, filepath.Join(ui.MainFilePath, "index.html"))
}
