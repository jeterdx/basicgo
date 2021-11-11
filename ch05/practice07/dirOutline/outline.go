package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// forEachNode は n から始まるツリー内の個々のノード x に対して
// 関数pre(x) と post(x) を呼び出します。 その二つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され （前順:preorder）、
// post は子ノードを訪れた後に呼び出されます （後順:postorder）。

func main() {
	doc, err := html.Parse(os.Stdin) //html.Parseの返り値はParse Tree
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

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

var depth int

func startElement(n *html.Node) {
	var keyList []string //属性のkeyを入れておくlistを用意
	if n.Type == html.ElementNode {
		for _, a := range n.Attr { //nodeのkeyを格納する
			keyList = append(keyList, a.Key)
		}
		if keyList == nil { //属性がない場合
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		} else { //属性がある場合
			fmt.Printf("%*s<%s ", depth*2, "", n.Data)
			for i, v := range keyList {
				if i+1 == len(keyList) { //list内で最後のkeyだったら閉じる
					fmt.Printf("%s='...'>\n", v)
				} else { //まだkeyが続く場合はkeyを出力するだけ
					fmt.Printf("%s='...' ", v)
				}
			}
			depth++
		}
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
