package popcount

// 1桁づつシフトしていって一つ一つのbitが1の時にカウンターを１増やす、64回繰り返す。
func PopCount(x uint64) int {
	var counter int
	for i := 0; i < 64; i++ {
		if refbit(x, 0) == 1 {
			counter++
		}
		x = x >> 1
	}
	return counter
}

func refbit(i uint64, b uint) uint64 {
	return (i >> b) & 1
}
