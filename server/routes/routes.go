package routes

import (
	"net/http"
	"travel-accommodation/db"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// home page
	routeRoot := http.HandlerFunc(rootHandler)
	mux.Handle("/", rootMiddleware(routeRoot))

	// db route
	routeDb := http.HandlerFunc(dbHandler)
	mux.Handle("/db", routeDb)

	return mux
}

// root route handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hello from home path"))
}

// db route handler
func dbHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	db.InsertUserData("Oscar", "Diaz", "mail@mail.com", 4248275260)
	w.Write([]byte("Query execute successfully"))
}
