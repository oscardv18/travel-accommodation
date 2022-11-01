package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-accommodation/server/routes"
)

func main() {
	fmt.Println("Starting Server at port 8080")

	if err := http.ListenAndServe(":8080", routes.Router()); err != nil {
		log.Fatal(err)
	}
}
