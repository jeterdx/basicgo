package popcountA

// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) //bit数を使いたいのでindexを利用する、ここで初期値を設定している。中身は0-255までの数を二進数で表した時の1の数を格納したもの。なぜこの式でできるかはわからない。
	}
}

// PopCount は x のポピュレーションカウント （1 が設定されているビット数） を返します。
func PopCount(x uint64) int {
	var counter int
	for i := 0; i < 8; i++ {
		counter += int(pc[byte(x>>(i*8))]) //8回足し算していたものを8回ループで回す。
	}
	return counter
}
