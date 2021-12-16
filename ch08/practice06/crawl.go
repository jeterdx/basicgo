package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	var depth int
	flag.IntVar(&depth, "depth", 0, "depth to crawl")
	flag.Parse()
	fmt.Println(depth)

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[2:]
	}()

	for i := 0; i < 20; i++ { //最大20個までのgoroutineを並列で並べる
		go func() {
			//mainのgoroutineが進んで、unseenlinksが追加されるまでここでストップされる。下にSleepを挟むとわかりやすい。
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
					depth-- //クロールと追加が終わったタイミングでdepthを減らす
					if depth == 0 {
						os.Exit(0) //プログラムを終了する
					}
				}()
			}
		}()
	}

	// メインゴルーチンは worklist の項目の重複をなくし、
	// 未探索の項目をクローラへ送る。
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] { //まだmap存在していなかった場合、
				seen[link] = true   //ステータスをtrueに変えて、
				unseenLinks <- link //まだクロールしていないリストに追加する
			}
		}
	}
}

// token は、 20個の並行なリクエストという限界を
// 強制するために使われる計数セマフォです。
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを解放
	if err != nil {
		log.Print(err)
	}
	return list
}
