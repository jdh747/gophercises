package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%v: %v", msg, err))
	}
}

func main() {
	file, err := os.Open("./ex.json")
	check("Error opening file", err)
	defer file.Close()

	story := make(map[string]Arc)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&story)
	check("Error decoding json", err)

	fmt.Println(story)
}

type Arc struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}
