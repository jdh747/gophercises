package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//TODO: pull file from cmdline arg
	file, err := os.Open("tests/ex4.html")
	check("Error opening file", err)

	node, err := html.Parse(file)
	check("Error parsing HTML", err)

	ParseTree(node, "a", func(aTagNode *html.Node) {
		for _, tag := range aTagNode.Attr {
			if tag.Key == "href" {
				fmt.Printf("Link {\n\tHref: '%v',\n\tText: '%v'\n}\n",
					tag.Val,
					strings.TrimSpace(aTagNode.FirstChild.Data))
			}
		}
	})
}

// ParseTree does something
func ParseTree(node *html.Node, nodeType string, printFunc func(*html.Node)) {
	if node.Type == html.ElementNode && node.Data == nodeType {
		printFunc(node)
	}

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
