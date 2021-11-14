package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

//与えられたURLのディレクトリを全てクロールし、ページのソースを複製して保存していく。

func breadthFirst(f func(item string) []string, worklist []string) {
	fmt.Println(worklist)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			fmt.Println(item)
			if !seen[item] { //defalutでfalseなので初見は入る。そしてtrueに変更。
				seen[item] = true
				copyPages(item)                         //URLを引数に関数を呼び出す。
				worklist = append(worklist, f(item)...) //ここでcrawlを呼んで結果をworklistにappendする。返ってくるのは、渡したURLのHTML内にあるリンクたち。
			}
		}
	}
}

func copyPages(url string) {
	//URLからbodyを取得
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copyPages: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "copyPages: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	/*URL・ファイル名・ディレクトリ名の操作がうまくできず断念
	//同じドメインかどうかの確認、ディレクトリを階層構造で作成、ページ内容をコピーしたファイルを作成
	if strings.HasPrefix(url, string(os.Args[1])) { //同じドメインかどうか確認
		err := os.MkdirAll(url, 0755)
		check(err)

		os.Create()
	}

}
*/

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url) //URLを受け取ってそのHTMLをParseし、その中にあるリンクのリストを返す。
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	err := os.Mkdir("copied", 0755)
	check(err)
	breadthFirst(crawl, os.Args[1:]) //URLは1つまでしか受け付けない
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
