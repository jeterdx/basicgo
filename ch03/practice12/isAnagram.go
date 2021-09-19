package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("listen", "silent"))
}

//アナグラムは同じ文字が同じ数存在するかを判定すること
//文字列を先頭から比較して、同じ文字が存在した場合、その文字同士をピリオドに置き換える、という処理を繰り返して最後に2つのstringがイコールになっていたらtrue、って風に描いてみる
func isAnagram(s1 string, s2 string) bool {
	for i1, v1 := range s1 {
		for i2, v2 := range s2 {
			if string(v1) == "" || string(v2) == "" {
				return false
			}
			if string(v1) == string(v2) {
				s1 = strings.Replace(s1, string(s1[i1]), ".", 1)
				s2 = strings.Replace(s2, string(s2[i2]), ".", 1)
			}
		}
	}
	fmt.Println(s1)
	fmt.Println(s2)
	return s1 == s2
}
