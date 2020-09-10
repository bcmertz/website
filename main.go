package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fileServer := http.FileServer(http.Dir("src"))
	http.Handle("/", fileServer)

	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
