package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func main() {
	c1 := sha256.Sum256([]byte("werqf"))
	c2 := sha256.Sum256([]byte("r"))
	fmt.Println("The number of different bits is ...")
	fmt.Println(diffBitHash(c1, c2))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) //bit数を使いたいのでindexを利用するだけ
	}
}

//それぞれのbyte配列のxorをとった結果のpopcountを変数diffに足していって値を返す。
func diffBitHash(c1 [32]byte, c2 [32]byte) int {
	var diff int
	for i := 0; i < 32; i++ {
		diff += int(pc[c1[i]^c2[i]])
	}
	return diff
}
