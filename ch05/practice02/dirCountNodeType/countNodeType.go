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

	result := make(map[string]int)
	for _, nodeType := range countNodeType(nil, doc) { //stringのリストから同じ要素の数をカウントして表記する処理
		result[nodeType]++
	}
	for k, v := range result {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func countNodeType(lists []string, n *html.Node) []string {
	if n.Type == html.ElementNode { // <html></html>から順にtypeを判別し、ElementNodeに該当するかを判定
		lists = append(lists, n.Data) //　ElementNodeであればそのタグ名をlistsに追加
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling { //今見ているnodeに子nodeがnilになるまで、子nodeをvisitに渡して再帰する。そして隣接するnodeをcにセットして、tree構造を下にたどっていく。
		lists = countNodeType(lists, c)
	}
	return lists
}
