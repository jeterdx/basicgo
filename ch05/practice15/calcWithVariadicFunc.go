package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(min(1))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(max())
	fmt.Println(min(-1, 2, 3, 4, 5))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("Stopped eintire program: Max func needs at least one value as an argument.")
		os.Exit(1)
		return 0
	}
	maxVal := -9223372036854775807 //intのとりうる最小値
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}

func min(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("Stopped eintire program: Min func needs at least one value as an argument.")
		os.Exit(1)
		return 0
	} else {
		minVal := 9223372036854775807 //intのとりうる最大値
		for _, val := range vals {
			if minVal > val {
				minVal = val
			}
		}
		return minVal
	}
}
