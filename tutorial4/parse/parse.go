package parse

import (
	"fmt"
	"strings"
	"os"

	"golang.org/x/net/html"
)

// ATags does something
func ATags(file *os.File) {
	node, err := html.Parse(file)
	Check("Error parsing HTML", err)

	parseTree(node, "a", func(aTagNode *html.Node) {
		for _, tag := range aTagNode.Attr {
			if tag.Key == "href" {
				fmt.Printf("Link {\n\tHref: '%v',\n\tText: '%v'\n}\n",
					tag.Val,
					strings.TrimSpace(aTagNode.FirstChild.Data))
			}
		}
	})
}

func parseTree(node *html.Node, nodeType string, printFunc func(*html.Node)) {
	if node.Type == html.ElementNode && node.Data == nodeType {
		printFunc(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		parseTree(child, nodeType, printFunc)
	}
}

// Check does something
func Check(msg string, err error) {
	if err != nil {
		fmt.Printf("%v: %v", msg, err)
		panic(err)
	}
}
