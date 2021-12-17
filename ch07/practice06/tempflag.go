package main

import (
	"basicgo/ch07/practice06/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 10.0, "the temperature") //ここの内部処理でflag.CommandLine.Varが走っているのでmain関数内でparseができる。

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
