package main

import (
	"os"

	"./parse"
)

func main() {
	file, err := os.Open(os.Args[1])
	parse.Check("Error opening file", err)

	parse.ATags(file)
}
