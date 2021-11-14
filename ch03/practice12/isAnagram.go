package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("aaaavvaav", "vvvaaaaaa")
	fmt.Println(isAnagram("aaaavvaav", "vvvaaaaaa"))

	fmt.Println("aaaa", "vvv")
	fmt.Println(isAnagram("aaaa", "vvv"))

	//日本語はできなかった。stringのままやろうとしているのがよろしくないポイ。
}

func isAnagram(s1 string, s2 string) bool {
	n := utf8.RuneCountInString(s1) //一つ目の文字列の長さをnに格納
	counter := 0                    //カウンターをゼロにセット
label1:
	for i1, v1 := range s1 {
		counter++ //1つ目の引数を全部回し切るためにカウンターを増やしていく
		for i2, v2 := range s2 {
			if (string(v1) == string(v2)) && !(string(v1) == ".") { //ピリオドに置換されていたケースを除き、等しい文字列の時に文字列変換処理をする
				s1 = strings.Replace(s1, string(s1[i1]), ".", 1)
				s2 = strings.Replace(s2, string(s2[i2]), ".", 1)
				break label1 //2重ループを抜けることで、ピリオドへの変換を反映しない状態での for文を一回抜ける。
			}
		}
	}
	if counter == n { //カウンターが文字列の分回りきったら、returnで結果を判定。
		return s1 == s2
	} else {
		return isAnagram(s1, s2) //更新済みのs1/s2で再帰的に関数を呼び出す。
	}
}
