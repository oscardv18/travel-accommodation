package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-accommodation/server/routes"
)

func main() {
	fmt.Println("Starting Server at port 8080")
	// routeRoot := http.HandlerFunc(routes.RootHandler)
	// http.Handle("/", routes.RootMiddleware(routeRoot))

	if err := http.ListenAndServe(":8080", routes.Router()); err != nil {
		log.Fatal(err)
	}
}
