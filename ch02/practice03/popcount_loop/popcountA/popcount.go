package popcountA

import (
	"fmt"
)

// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) //bit数を使いたいのでindexを利用するだけ
	}
}

// PopCount は x のポピュレーションカウント （1 が設定されているビット数） を返します。
func PopCount(x uint64) int {
	fmt.Println(pc)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
