package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//Start measuring time
	now := time.Now()

	//Main part of program
	fmt.Println(strings.Join(os.Args[1:], " "))

	//Print the time taken for the execution
	fmt.Printf("Result : %vms\n", time.Since(now).Milliseconds())

}
