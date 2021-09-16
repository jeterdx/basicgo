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
	return int(pc[byte(x>>(0*8))] + //255まではここで対応できる。
		pc[byte(x>>(1*8))] + //それ以上に桁数がある場合は、8bitシフトさせて次の8bitの1の数をカウントする。byteなので8bitづつ。
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
