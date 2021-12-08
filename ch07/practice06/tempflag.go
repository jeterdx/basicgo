package main

import (
	"basicgo/ch07/practice06/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature") //ここの内部処理でflag.CommandLine.Varが走っているのでmain関数内でparseができる。

func main() {
	flag.Parse() //このParse関数を呼ぶことでなぜtempconvパッケージに定義しているSet関数が呼ばれているのかが理解できない。
	fmt.Println(*temp)
}
