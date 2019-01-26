package urlshort

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // This import registers pq as a driver for postgres SQL databases

	yaml "gopkg.in/yaml.v2"
)

// MapHandler ...
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if redirect, ok := pathsToUrls[r.URL.RequestURI()]; ok {
			http.Redirect(w, r, redirect, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}

// DbHandler ...
func DbHandler(fallback http.Handler) (http.HandlerFunc, error) {
	connStr := "user=golang dbname=golang_test_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var redirect string
		err = db.QueryRow("SELECT url FROM url_shortner WHERE path = $1", r.URL.RequestURI()).Scan(&redirect)
		if err == sql.ErrNoRows {
			fallback.ServeHTTP(w, r)
		} else if err != nil {
			log.Fatal(err)
		} else {
			http.Redirect(w, r, redirect, http.StatusMovedPermanently)
		}
	}), nil
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
