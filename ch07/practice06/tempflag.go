package main

import (
	"flag"
	"fmt"

	"basicgo/ch07/practice06/tempconv/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
