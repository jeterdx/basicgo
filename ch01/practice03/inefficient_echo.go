package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	//Start measuring time
	now := time.Now()

	//Main part of program
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//Print the time taken for the execution
	fmt.Printf("Result : %vms\n", time.Since(now).Milliseconds())
}
