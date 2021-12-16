package main // pc[i] は i のポピュレーションカウントです。

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; 10000 > i; i++ {
		go func() {
			fmt.Println(PopCount(uint64(i)))
		}()
	}
}

var pc [256]byte

func initPopcount() { //initのままだとDoの引数にできないから変更
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var initPopcountOnce sync.Once //sync.Onceを用意

// PopCount は x のポピュレーションカウント （1 が設定されているビット数） を返します。
func PopCount(x uint64) int {
	initPopcountOnce.Do(initPopcount) //Popcountが呼び出された最初の一回だけロックをかける。しっかりと変数がinitPopcountで初期化されることを待ってから、goroutine達を実行する。
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
