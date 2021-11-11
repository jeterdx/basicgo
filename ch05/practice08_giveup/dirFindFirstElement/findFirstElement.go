package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v\n", err)
		os.Exit(1)
	}
	node := ElementByID(doc, "script")
	if node != nil {
		fmt.Printf("<%s> was found", node.Data)
	} else {
		fmt.Println("No node with the id was found")
	}
}

//forEashNodeの修正の仕方が分からない。
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//ElementByIDも最初のnodeしか精査できず、forEachNodeの組み込み方が分からない。
func ElementByID(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode {
		for _, a := range doc.Attr {
			if a.Key == "id" && a.Val == id {
				return doc
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(c, id)
	}
	return nil
}
