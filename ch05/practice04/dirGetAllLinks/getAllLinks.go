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
	for _, link := range visit(nil, doc) { //visitの返り値の<a href>に該当するvalueの[]stirngをrangeで回してる。
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	//再帰的に呼び出すので最後までnodeを精査した段階でエラーが生じないようにreturnしておく
	if n == nil {
		return links
	}
	if (n.Type == html.ElementNode) && (n.Data == "a" || n.Data == "image" || n.Data == "link" || n.Data == "script") {
		for _, a := range n.Attr { //hrefは属性、attribute
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	//再帰呼び出しとして、子node、隣接nodeをそれぞれvisitに渡す
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
