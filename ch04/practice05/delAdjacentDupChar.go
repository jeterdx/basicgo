package main

import "fmt"

func main() {
	s := []string{"a", "a", "a", "c", "d", "e", "e", "e"}
	fmt.Println(delAdjacentDupChar(s))
}

//重複しない文字列を見つけるまでloopをスキップして、見つけたらn+1の文字列をi+xの文字列に置き換える
func delAdjacentDupChar(s []string) []string {
	n := 0
	for i := 0; i < len(s)-1; i++ {
		if s[n] != s[i+1] {
			s[n+1] = s[i+1]
			n = n + 1
		}
	}
	return s[:n+1]
}

// func remove(slice []string, i int) []string {
// 	copy(slice[i:], slice[i+1:])
// 	return slice[:len(slice)-1]
// }

// func equal(x, y []string) bool {
// 	if len(x) != len(y) {
// 		return false  }
// 	for i := range x {
// 		if x[i] != y[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
