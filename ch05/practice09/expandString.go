package main

import (
	"fmt"
	"strings"
)

func main() {
	var info struct {
		s string
		f func(string) string
	}

	info.s = "ABCD  &AN$BC $$aaaa  D"
	info.f = strings.ToLower //一つ目の$以降を小文字にする

	fmt.Println(expand(info.s, info.f))
}

func expand(s string, f func(string) string) string {
	index := strings.Index(s, "$")
	if index == -1 { //$がなかったらそのままsを返す
		return s
	} else {
		s = strings.Replace(s, s[index:], f(s[index:]), -1)
		return s
	}
}
