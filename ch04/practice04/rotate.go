package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a[:], 3))
}

//スライスの0番目の要素をスライスの最後にappendする処理をt回繰り返す。
func rotate(s []int, t int) []int {
	for ; t > 0; t-- {
		first := s[0]            //最初の要素
		s = append(s[1:], first) //最初の要素を除く配列の最後にfirstをappendする。
	}
	return s
}
