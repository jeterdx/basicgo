package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//Parseしてポート番号を受け取る
	var portNum int
	flag.IntVar(&portNum, "port", 0, "port number")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", portNum))
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) //cはio.Writer型でもあるから引数に取れる。なぜならnetパッケージのconnインタフェースがWriteメソッドを実装してるから。WriteStringを使ってるから普通のStringを書き込めている。
		if err != nil {
			return // 例: クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
