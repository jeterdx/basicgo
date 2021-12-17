package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m2s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("5m24s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2045, length("4m24s")},
}

//文字列で受け取った時間をtime.Duration型へParseして返す関数
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//表を出力している
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 列幅を計算して表を印字する
}

//以下はより柔軟なソートを実現するための構造体の定義
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool //*Trackをレシーバとして受け取って使えるようにしている無名関数less。返り値はbool。
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	//カスタムソートを使って、Keyをコマンドライン引数から受け付ける。keyListとしてforで回して一致するタイミングで比較を実施する。
	keyList := os.Args[1:]
	sort.Sort(customSort{tracks, func(x, y *Track) bool { // customSortはtracksと無名関数を定義した構造体なのでそれを表している。無名関数の中身でSortの優先順位を表現している。
		for _, v := range keyList {
			if v == "Title" {
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			}
			if v == "Year" {
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			}
			if v == "Length" {
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			}
		}
		return false
	}})
	printTracks(tracks)

}
