package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin) //html.Parseの返り値はParse Tree
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, content := range visit(nil, doc) { //visitの返り値の<a href>に該当するvalueの[]stirngをrangeで回してる。
		fmt.Println(content)
	}

}

func visit(contents []string, n *html.Node) []string {
	if n.Type == html.TextNode { //nodeのタイプがテキストノードのDataをappendしていく
		contents = append(contents, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { //今見ているnodeに子nodeがnilになるまで、子nodeをvisitに渡して再帰する。そして隣接するnodeをcにセットして、tree構造を下にたどっていく。
		contents = visit(contents, c)
	}
	return contents
}
