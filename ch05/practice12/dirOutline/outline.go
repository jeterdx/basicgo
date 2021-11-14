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
	outline(doc)
}

// forEachNode は n から始まるツリー内の個々のノード x に対して
// 関数pre(x) と post(x) を呼び出します。 その二つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され （前順:preorder）、
// post は子ノードを訪れた後に呼び出されます （後順:postorder）。

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func outline(n *html.Node) { //outline関数を作って、forEachNodeに無名関数でstartElement/endElementをそれぞれ渡す。
	var depth int
	forEachNode(n, func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}, func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}) //forEachNodeの引数がここまで
}
