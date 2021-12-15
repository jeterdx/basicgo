package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例:接続が切れた
			continue
		}
		go handleConn(conn) // 一度に一つの接続を処理する
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c) //NewScannerの引数に取れるのはio.Reader型。os.Stdinもconnも満たしている。この一文でコンソールから標準入力を受け付ける。
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
