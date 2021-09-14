package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //チャネルは型じゃない、ゴルーチン間でデータの受け渡しを行うパイプのようなもの。そのデータの型がstringで定義されている。
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //ゴルーチン開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //チャンネルから受信
	}
	fmt.Println("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //chチャネルにエラーメッセージを送信
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
