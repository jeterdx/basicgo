package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte{'g', 'o', 'l', ' ', ' ', 'a', 'n', 'g'}
	fmt.Println(dupUnicodeSpaceToAsciiSpace(b))
}

//
func dupUnicodeSpaceToAsciiSpace(b []byte) []byte {
	n := 0
	for i := 0; i < len(b); i++ {
		if !(unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1]))) {
			//nの要素にASCII Spaceを入れ、n+1の要素（asciiスペースじゃない要素）を2つ目のascii spaceに入れる。
			b[n] = " " //byte型に対してascii spaceを入れるには。
			b[n+1] = b[i+1]
			n = n + 1
		}
	}
	return b[:n+1]
}
