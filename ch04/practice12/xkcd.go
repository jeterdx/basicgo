package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

func main() {
	type commicInfo struct {
		Num        int    `json:"num"`
		Year       string `json:"year"`
		Month      string `json:"month"`
		Day        string `json:"day"`
		Title      string `json:"title"`
		SafeTitle  string `json:"safe_title"`
		Link       string `json:"link"`
		News       string `json:"news"`
		Transcript string `json:"transcript"`
		Alt        string `json:"alt"`
		Img        string `json:"img"`
	}

	index := make(map[int]commicInfo)
	args := os.Args[1]

	for i := 1; 2529 > i; i++ { //めっちゃ時間かかるのでもっといいやり方がありそう。goroutineで並行処理。
		//URLを作る
		numStr := strconv.Itoa(i)
		endpoint := "https://xkcd.com/"
		u, err := url.Parse(endpoint)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
			os.Exit(1)
		}
		u.Path = path.Join(u.Path, numStr, "/info.0.json")

		//urlの先からresponseをもらってbodyだけ格納しておく
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close() // 自分で閉じる

		//bodyをjson形式の構造体に変換。
		structedJSON := []byte(body)
		var data commicInfo
		json.Unmarshal(structedJSON, &data)
		index[i] = data

		//コマンドライン引数で受け付けたタイトルと一致するものを出力
		if index[i].Title == string(args) {
			fmt.Println(index[i].Link)
			fmt.Println(index[i].Transcript)
		}
	}
	fmt.Println("---------------------------")
	fmt.Println("This is the end of index. If you cannot find resources you want, check the title you entered.")
}
