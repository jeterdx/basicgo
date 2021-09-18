package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345785439"))
}

func comma(s string) string {
	var buf bytes.Buffer

	l := len(s)
	for i, v := range s { //文字列をrangeで回すと、index/valueが返ってくる。rangeなのでこのvalueはrune。byte単位で取り出すのはsliceでアクセスした時。
		if (l-i)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v)) //Fprintfは第一引数が出力先（io.Writer）で、%dが10進数形式をStringで表示する、vがvalueでここではrune。
	}
	return buf.String()
}
