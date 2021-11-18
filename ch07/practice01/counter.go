package main

import (
	"bufio"
	"fmt"
)

func main() {
	var cw countWord
	cw.Write([]byte("hello my name is ..."))
	fmt.Println(cw)

	var cl countLine
	cl.Write([]byte("hello my name is ...\n. hello my name is ...\n. hello my name is ...\n. hello my name is ...\n. hello my name is ...\n. hello my name is ...\n."))
	fmt.Println(cl)
}

type countWord int

func (c *countWord) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanWords(p, true) //最初のIsSpaceがtrueになるまでの[]byteをtokenとして返し、そのtokenの次のindexがadvaceとして返される。
		if err != nil {
			return 0, err
		}
		if token != nil { //tokenが空になるまでカウントする。
			*c++
		}
		//fmt.Println(string(token))
		//fmt.Println(string(p[advance:]))
		p = p[advance:] //[]byteをtoken部分を除外して更新する
	}
	return n, nil
}

type countLine int

func (c *countLine) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanLines(p, true)
		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}
	return n, nil
}
