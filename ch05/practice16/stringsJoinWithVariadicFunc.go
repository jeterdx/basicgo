package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Join("-", "a", "あ", "c"))
	fmt.Println(Join("-"))
}

func Join(sep string, elems ...string) string {
	if len(elems) < 2 {
		fmt.Println("Need at least two strings to join.")
		os.Exit(1)
		return "e" //到達しないのでなんでも良い
	} else {
		result := ""
		for i, v := range elems {
			if i == 0 { //1つ目の要素の時はsepを追加したくない
				result = v
			} else {
				result += sep + v
			}
		}
		return result
	}
}
