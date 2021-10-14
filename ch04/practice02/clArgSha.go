package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

//なんだかコマンドライン引数からの入力を受け取るタイミングとかがflagが上手く機能している気がしない。。
func main() {

	//flagの使用
	argument := flag.String("type", "sha256", "sha256-384-512 can be used.")
	flag.Parse()

	//stdinを読み込む
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//オプションの判定とそれに応じた処理
		if !(*argument == "sha256" || *argument == "sha384" || *argument == "sha512") {
			fmt.Fprintf(os.Stderr, "invalid argument (input sha256/384/512): %s\n", *argument)
			os.Exit(1)
		} else if *argument == "sha256" {
			input := scanner.Bytes()
			result := sha256.Sum256(input)
			fmt.Printf("%x", result)

		} else if *argument == "sha384" {
			input := scanner.Bytes()
			result := sha512.Sum384(input)
			fmt.Printf("%x", result)

		} else if *argument == "sha512" {
			input := scanner.Bytes()
			result := sha512.Sum512(input)
			fmt.Printf("%x", result)
		}
	}
}
