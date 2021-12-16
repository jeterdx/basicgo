package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	input := bufio.NewScanner(c) //NewScannerの引数に取れるのはio.Reader型。os.Stdinもconnも満たしている。この一文でコンソールから標準入力を受け付ける。
	for input.Scan() {
		wg.Add(1)   //gorouineの外に書かないと、Closerの呼び出しの前に必ずAddが実行されているかどうかが保証できない。
		go func() { //Doneを書きたいので無名関数の中にechoを入れる。
			echo(c, input.Text(), 1*time.Second)
			defer wg.Done()
		}()
	}
	//以下Closer
	go func() {
		wg.Wait() //wgが0になるまでgoroutineの実行を停止
		tcpConn, ok := c.(*net.TCPConn)
		if !ok {
			log.Fatal("cast to TCPConn did not succeed")
		}
		tcpConn.CloseWrite() //カウンターが0、全てのechoが終了したら書き込みのコネクションをクローズする。
	}()
}
