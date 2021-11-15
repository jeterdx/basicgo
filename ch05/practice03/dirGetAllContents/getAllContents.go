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
	//fmt.Println(visit(nil, doc))

}

func visit(contents []string, n *html.Node) []string {
	if n.Type == html.TextNode { //nodeのタイプがテキストノードのDataをappendして
		contents = append(contents, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { //今見ているnodeに子nodeがnilになるまで、子nodeをvisitに渡して再帰する。そして隣接するnodeをcにセットして、tree構造を下にたどっていく。
		contents = visit(contents, c)
	}
	//改行だけの要素とかが邪魔なので取り除いて新しいリストを返す
	var contentsDumped []string
	for _, content := range contents {
		for _, r := range content {
			if r != 10 {
				contentsDumped = append(contentsDumped, content)
				break
			}
		}
	}
	return contentsDumped
}
