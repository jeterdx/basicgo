package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() { //mainにもgorouineが実行されてる。イメ-ジは22行目まで実行してから、go文を見つけてもう一つgoroutineを生成。定義されてる無名関数を実行しつつ、かつ並行して27行目に飛ぶ。22行目から25行目が完了すると、29行目を実行できる、って流れ。
	conn, err := net.Dial("tcp", "localhost:8000") //Dialでコネクションを確立する。
	if err != nil {
		log.Fatal(err)
	}
	//connのうち、*net.TCPConn型を使えるように型アサーションする
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("cast to TCPConn did not succeed")
	}

	done := make(chan struct{})
	go func() { //無名関数の定義とgoroutineでの実行を同時にしている。go func(){関数の処理内容}(実際に無名関数実行で渡す引数)、という構造になっている。ここでは引数を取らない関数なので26行目に()がある。
		io.Copy(os.Stdout, conn) // 標準出力に、connをコピーする。io.Copyはconnに読み込めるデータが来るまで実行されない。一度実行されるとEOFになるまでos.Stdoutにcopyする。connにデータが来るのは、27行目でmustCopyが実行されてから。
		log.Println("done")
		done <- struct{}{} // メイン関数のゴルーチンへ無名関数の処理が終わったことを通知、strcut{}{}は空の構造体を宣言＆値取りしてる。このように空の値を渡すのは、intやboolだと同期以外の目的になりうる値が入る可能性があるから。
	}()
	mustCopy(conn, os.Stdin) //標準入力で受け取ったものを、connに書き込む
	tcpConn.CloseWrite()     //ここをtcpConnのCloseWriteでクライアントからの書き込みだけCloseし、サーバ側からの送信は続けるようにする。
	<-done                   // バックグラウンドのゴルーチン、つまり、22行目で定義されてる無名関数のやつが完了するまで、main関数を終わらせるのを待つ
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
