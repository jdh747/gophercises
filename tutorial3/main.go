package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
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

	story := make(map[string]Arc)
	err = json.NewDecoder(file).Decode(&story)
	check("Error decoding json", err)

	template, err := template.ParseFiles("template.html")
	check("Error parsing template", err)
	err = template.Execute(os.Stdout, story["intro"])
	check("Error executing template", err)
}

type Arc struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}
