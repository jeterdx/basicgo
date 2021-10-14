package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //チャネルは型じゃない、ゴルーチン間でデータの受け渡しを行うパイプのようなもの。そのデータの型がstringで定義されている。
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //ゴルーチン開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //チャンネルから受信したデータを出力している
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds()) //プログラム全体の経過時間をプリントしている
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //ここがよくわからない。
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //resp.Bodymをioutil.Discardにコピーしている、ioutil.Discardは書き込まれたデータ全てを破棄している。返り値はコピーしたバイト数とエラー。i
	resp.Body.Close()                                 // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()                  //そのゴルーチンが始まってからの経過時間を格納してる
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) //データをchチャネルに要約して送っている。

}
