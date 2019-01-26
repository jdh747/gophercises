package urlshort

import (
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler ...
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if redirect, ok := pathsToUrls[r.URL.RequestURI()]; ok {
			http.Redirect(w, r, redirect, http.StatusMovedPermanently)
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler ...
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return MapHandler(yamlToMap(yml), fallback), nil
}

func yamlToMap(yml []byte) map[string]string {
	var urlEntries []map[string]string
	err := yaml.Unmarshal(yml, &urlEntries)
	if err != nil {
		log.Fatal(err)
	}

	urls := make(map[string]string)
	for _, entry := range urlEntries {
		urls[entry["path"]] = entry["url"]
	}
	return urls
}