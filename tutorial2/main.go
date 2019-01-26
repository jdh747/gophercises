package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./urlshort"
)

func main() {
	// Default handler
	mux := defaultMux()

	// Map handler
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// YAML handler
	yamlFile := flag.String("yaml", "", "filename of yaml file with list of [path, url] entries used in url mapping")
	flag.Parse()

	var yaml []byte
	var err error
	if *yamlFile != "" {
		yaml, err = ioutil.ReadFile(*yamlFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Serve handlers
	port := ":8082"
	fmt.Println("Starting the server on ", port)
	if err = http.ListenAndServe(port, yamlHandler); err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
