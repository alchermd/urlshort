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
	yaml := `
- path: /netflix
  url: https://netflix.com/
- path: /google
  url: https://google.com/
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	log.Print("Starting server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}
