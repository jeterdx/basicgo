package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsFloat("1234124145785439.3525525"))
}

func IsFloat(s string) string {
	sAfterP := ""
	sBeforeP := s
	for i, v := range s {
		if string(v) == "." {
			sAfterP = s[i:]
			sBeforeP = s[:i-1]
		}
	}
	return comma(sBeforeP) + sAfterP
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
