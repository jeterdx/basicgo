package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(os.Stdout, resp.Body) //io.Copy関数でbodyフィールドを標準出力にコピー。返り値はコピーしたデータのbyte数とエラー型
		resp.Body.Close()

		if err != nil { //エラー処理は変わらず
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\nThe volume of body is %d bytes", n) //bodyが標準出力で出された後にバイト数もプリントしておく
	}
}