package main

import (
	"fmt"
	"sort"
	"unicode"
)

type runeList []rune

//runeListがLen、Less、Swapのメソッドを持ってsort.Interfaceを満たすようにする
func (r runeList) Len() int {
	return len(r)
}

func (r runeList) Less(i, j int) bool {
	//アルファベットの時は小文字で比較するようにする
	if unicode.IsLetter(r[i]) {
		r[i] = unicode.ToLower(r[i])
	}
	if unicode.IsLetter(r[j]) {
		r[j] = unicode.ToLower(r[j])
	}
	return r[i] < r[j]
}

func (r runeList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func isPalindrome(s sort.Interface) bool {
	//回文かどうかを判断する処理をここに書く
	length := s.Len() //len(s)だとエラーが出る。
	for i := 0; i < length/2; i++ {
		j := length - 1 - i
		if !s.Less(i, j) && !s.Less(j, i) { //両方の結果がfalse、つまりどっちから見てもLessじゃない、同じだったら、処理を続ける
			continue
		}
		return false //ループを抜けちゃったらfalse
	}
	return true //for文の条件式を満たさなくなるまで回せたら回文なのでtrue
}

func main() {
	text := "bababababbbbbb"
	var rl = runeList(text)
	if isPalindrome(rl) {
		fmt.Println(text + " is palindrome")
	} else {
		fmt.Println(text + " is not palindrome")
	}
}
