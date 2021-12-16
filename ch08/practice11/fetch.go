//複数のリクエストを平行で処理して一番速かったレスポンスのみを返します。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	result := mirroredFetch()
	fmt.Println(string(result))
}

func mirroredFetch() []byte {
	responses := make(chan []byte)
	go func() { responses <- fetch("https://google.com") }() //go.devといい勝負。
	go func() { responses <- fetch("https://go.dev/doc/") }()
	go func() { responses <- fetch("https://youtube.com") }()
	//close(responses) closeしたらゼロ値になっちゃう。mirroedQuesryに従ってというのが理解できていない。
	return <-responses // return the quickest response
}

func fetch(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return b
}
