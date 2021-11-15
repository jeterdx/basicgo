package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin) //html.Parseの返り値はParse Tree
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v\n", err)
		os.Exit(1)
	}

	images := ElementsByTagName(doc, "img")
	fmt.Printf("Here is the number of img tag :%d\n", len(images))

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Printf("Here is the number of h tag :%d", len(headings))
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	//nodelistを作って、受け取ったnodeがnameのlistと等しい場合にappendしていく
	var nodeList []*html.Node
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				nodeList = append(nodeList, doc)
			}
		}
	}
	//nodeを子・隣接と探索していく。可変引き数なので name...を引数に与える。返ってきたnodelistのnodeを1つずつnodelistに追加する。
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, node := range ElementsByTagName(c, name...) {
			nodeList = append(nodeList, node)
		}
	}
	return nodeList
}
