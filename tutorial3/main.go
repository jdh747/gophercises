package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func check(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%v: %v", msg, err))
	}
}

func main() {
	//todo: package everything nicely
	file, err := os.Open("ex.json")
	check("Error opening file", err)
	defer file.Close()

	story := make(map[string]arc)
	err = json.NewDecoder(file).Decode(&story)
	check("Error decoding json", err)

	template, err := template.ParseFiles("template.html")
	check("Error parsing template", err)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc := strings.Split(r.RequestURI, "/")[1]
		if arc == "" {
			arc = "intro"
		}

		err = template.Execute(w, story[arc])
		check("Error executing template", err)
	})

	http.ListenAndServe(":8080", mux)
}

type arc struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}