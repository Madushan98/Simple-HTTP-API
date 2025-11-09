package main

import (
	"log"
	"net/http"

	"simple-http-api/api/handler"
)

func main() {
	http.HandleFunc("/hello-world", handler.HelloWorldHandler)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
