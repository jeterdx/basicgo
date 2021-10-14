package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // Unicode文字の数、keyがruneでコードポイント、valueが出現回数
	isType := ""
	invalid := 0 //不正な UTF-8文字の数

	in := bufio.NewReader(os.Stdin)
	for { //EOFでbreakなので条件式などは描かない、そのほかのerrは2つ目のifで処理
		r, n, err := in.ReadRune() // rune, nbytes, error を返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { //invalidな文字の判定
			invalid++
			continue
		}
		//こっからジャンルごとに分類
		if unicode.IsLetter(r) {
			isType = "Letter"
			counts[isType]++ //countマップでisTypeをkeyにしてカウンタを1つ増やす。
		} else if unicode.IsNumber(r) {
			isType = "Number"
			counts[isType]++ //countマップでisTypeをkeyにしてカウンタを1つ増やす。
		} else if unicode.IsSpace(r) {
			isType = "Space"
			counts[isType]++ //countマップでisTypeをkeyにしてカウンタを1つ増やす。
		} else {
			isType = "Others"
			counts[isType]++ //countマップでisTypeをkeyにしてカウンタを1つ増やす。
		}
	}

	fmt.Printf("UnicodeType\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
