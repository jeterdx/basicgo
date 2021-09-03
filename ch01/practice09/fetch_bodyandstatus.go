package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //httpパッケージのGetメソッドがurlを引数として、2つの返り値を返している。1つ目がhttpレスポンス全体、2つ目がエラーがあった時のエラー。レスポンス全体は構造体型。
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) //標準エラー出力からエラーメッセージを出力する。例えば、no such hostとかがあった。
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body) //resp構造体のフィールドの1つ、Bodyフィールド全体を読み込む。
		resp.Body.Close()                   //変数bに格納したら資源節約のためストリームを閉じる。
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//HTTP Statusも同じ容量で抽出
		s := resp.Status //resp構造体のフィールドの1つ、Bodyフィールド全体を読み込む。
		fmt.Printf("%s", b)
		fmt.Printf("\n%s", s)
	}
}
