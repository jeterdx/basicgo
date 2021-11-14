package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'} //これもpractice06に同じく、byte列じゃない。
	fmt.Println(string(reverse(b)))
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
