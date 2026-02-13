package main

import (
	"log"
	"net/http"
)

func main() {
	const root = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(root)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s at http://localhost%s", root, server.Addr)
	log.Fatal(server.ListenAndServe())

}
