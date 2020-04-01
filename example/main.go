package main

import (
	"flag"
	"fmt"
	"github.com/alchermd/urlshort"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	yamlFileName := flag.String("yaml", "urls.yaml", "the YAML file that contains URL definitions")
	flag.Parse()

	yaml, err := ioutil.ReadFile(*yamlFileName)
	if err != nil {
		fmt.Println("Cound't open YAML file.")
		os.Exit(1)
	}

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/aws":   "https://aws.amazon.com/",
		"/gcp":   "https://cloud.google.com/",
		"/azure": "https://azure.microsoft.com/",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
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
