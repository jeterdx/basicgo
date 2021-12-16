package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
	confirmInput := make(chan struct{}) //標準入力から１バイトでも読み込むことができればconfirmImputチャネルにイベントを送信して以下のSelectの2番目に入れるようにする。それがないまま10秒経過すると出力とともにコネクションを閉じる。
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイトを読み込む
		confirmInput <- struct{}{}
	}()

	input := bufio.NewScanner(c) //NewScannerの引数に取れるのはio.Reader型。os.Stdinもconnも満たしている。この一文でコンソールから標準入力を受け付ける。
	select {
	case <-time.After(10 * time.Second): //なぜか10秒経ってもこのケースに入らない。
		fmt.Println("Connection closed since no input for 10sec.")
		c.Close()
	case <-confirmInput:
		for input.Scan() {
			fmt.Println("aaaa")
			echo(c, input.Text(), 1*time.Second)
		}
	}
}
