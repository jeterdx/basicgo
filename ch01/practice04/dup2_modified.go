package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)         //keyがstring型かつ値がint型のmap型のデータを作ってcountsに格納する
	filesources := make(map[string]string) //keyをlineとし、valueを使用したファイル群を文字列で結合したものを、新たにmap変数で作成する。1つのkeyで2つのvalueを格納できる変数はなさそうだったので。
	files := os.Args[1:]                   //コマンドライン引数をfiles変数に格納する、0は入力コマンド自身なので1から。
	for _, arg := range files {            //filesからファイルを取り出し、argに格納し、loopを回す
		f, err := os.Open(arg) //os.Openが返す、開いたファイルの中身と組み込みのerr型の値をそれぞれ格納する。
		if err != nil {        //os.Openが何かしらのエラー型の値を返していたら、プリント関数でエラー内容を表示する。
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg, filesources) //countline関数に、ファイルの中身の格納先のポインタ *os.Fileとcounts変数、現在のループで使用しているファイル名、それを格納するmap変数を渡して処理を実行させる。
		f.Close()
	}
	for line, n := range counts { //count変数の中身を表示する。
		if n > 1 { //valueが１より大きい、つまり、重複しているラインを表示する。
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", filesources[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, filesources map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] == 1 { //keyが1ということは初めてlineがカウントされたということ、その場合は以下のようにファイル名を追加する
			filesources[input.Text()] = filename
		} else if !(strings.Contains(filesources[input.Text()], filename)) { //keyが1じゃないとき、つまり、２以上の時には、同一ファイル名からの重複かを判断する必要がある。同一ファイル名が文字列として含まれない場合に、新しいファイル名を結合する。
			filesources[input.Text()] += "---" + filename //ファイル名をうまく結合する
		} else { //それ以外の条件時、keyが0、もしくは、ファイル名がすでに書き込まれている場合は何もしない
		}
	}
}
