package popcount

// 1桁づつシフトしていって一つ一つのbitが1の時にカウンターを１増やす、64回繰り返す。
func PopCount(x uint64) int {
	counter := 1
	if x == 0 {
		return 0
	} else {
		for i := 0; i < 64; i++ {
			x = x & (x - 1) //下一桁が1であれば1をクリアする
			if x > 0 {
				counter++
			}
		}
		return counter
	}
}
