// Netcat1 は読み込み専用の TCP クライアントです。

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	var tzs []string
	var paths []string

	for _, v := range args {
		eaqulIndex := strings.Index(v, "=")
		if eaqulIndex == -1 {
			fmt.Println("Please type valid option : [TIMEZONE]=localhost:[portNumber]")
		} else {
			paths = append(paths, v[eaqulIndex+1:])
			tzs = append(tzs, v[:eaqulIndex])
		}
	}

	fmt.Println(tzs)

	for _, v := range paths {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			log.Fatal(err)
		}
		//defer conn.Close()
		go mustCopy(os.Stdout, conn) //２つ目を標準出力した段階でコネクションが切れてしまう。
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
