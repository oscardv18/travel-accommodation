package main

import "net/http"

// root route handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 WELCOME"))
}
