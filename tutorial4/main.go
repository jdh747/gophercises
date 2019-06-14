package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	//TODO: pull file from cmdline arg
	file, err := os.Open("tests/ex1.html")
	check("Error opening file", err)

	node, err := html.Parse(file)
	check("Error parsing HTML", err)

	ParseTree(node, "a", func(aTagNode *html.Node) {
		fmt.Printf("Link{\n\tHref: '%v',\n\tText: '%v'\n}", aTagNode.Attr[0].Val, aTagNode.FirstChild.Data)
	})
}

// ParseTree does something
func ParseTree(node *html.Node, nodeType string, printFunc func(*html.Node)) {
	if node.Type == html.ElementNode && node.Data == nodeType {
		printFunc(node)
	}

	// Recurse through all children
	// TODO: could each child be handled by a goroutine?
	// TODO: might need a switch for dealing with various types of nodes
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ParseTree(child, nodeType, printFunc)
	}
}

func check(msg string, err error) {
	if err != nil {
		fmt.Printf("%v: %v", msg, err)
		panic(err)
	}
}
