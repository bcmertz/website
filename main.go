package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	log.Println("Parsing flags")
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("src"))
	http.Handle("/", fileServer)

	log.Printf("Starting server on port %v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
