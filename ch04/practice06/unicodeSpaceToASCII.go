package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte{'g', 'o', 'l', ' ', ' ', ' ', 'a', 'n', 'g', ' ', ' ', 'a'} //UTF-8でエンコードされた、なので、全角スペースなども対応できないといけない。これだとascii文字しか入っていないbyte[]を、になってしまう。
	fmt.Println(string(dupUnicodeSpaceToAsciiSpace(b)))                     //正しいやり方は、utf8.DecodeRuneを使い、何バイトのruneかを判定して正確に区切らないといけない。返り値がruneとサイズなのでそのサイズで区切ること。
}

func dupUnicodeSpaceToAsciiSpace(b []byte) []byte {
	for i := 0; i < len(b)-1; i++ {
		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1])) {
			b[i] = ' '                    //1つ目のスペースにasciiスペースを代入
			b = append(b[:i], b[i+1:]...) //1つ目のスペースまでの要素と2つ目のスペース以降の要素をくっつける
			i--                           //要素が減った分indexを1つ前にずらす
		}
	}
	return b
}
