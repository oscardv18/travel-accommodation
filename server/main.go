package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Server at port 8080")
	routeRoot := http.HandlerFunc(rootHandler)
	http.Handle("/", rootMiddleware(routeRoot))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
