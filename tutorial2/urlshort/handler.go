package urlshort

import (
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
	pathsToUrls, err := yamlToMap(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func yamlToMap(yml []byte) (map[string]string, error) {
	var urlEntries []map[string]string
	if err := yaml.Unmarshal(yml, &urlEntries); err != nil {
		return nil, err
	}

	urls := make(map[string]string)
	for _, entry := range urlEntries {
		urls[entry["path"]] = entry["url"]
	}
	return urls, nil
}
