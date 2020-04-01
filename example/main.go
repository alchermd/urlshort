package main

import (
	"fmt"
	"github.com/alchermd/urlshort"
	"log"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/aws":   "https://aws.amazon.com/",
		"/gcp":   "https://cloud.google.com/",
		"/azure": "https://azure.microsoft.com/",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	log.Print("Starting server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}
