package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	var (
		counts = make(map[string]int)
		input  = bufio.NewScanner(f)
	)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	//出力部分
	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
