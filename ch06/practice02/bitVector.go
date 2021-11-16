package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3, 4, 5, 64, 65, 128)
	fmt.Println(x.Len())
	fmt.Println(x.String())

	x.AddAll(300)
	fmt.Println(x.String())
}

func (s *IntSet) AddAll(xList ...int) {
	for _, v := range xList {
		word, bit := v/64, uint(v%64) //xの商と余剰をそれぞれwordとbitに入れる
		for word >= len(s.words) {    //最初はs.wordsは空。s.wordsのlenは0で、xが1だとwordは0なので条件を満たす
			s.words = append(s.words, 0) //0をリストに追加する。lenが1になる。なので１回でループを抜ける。65以上から2つ目の要素が追加される。1つの要素で64個まで整数を保管。
		}
		s.words[word] |= 1 << bit //保管した数値番目のbitを1にしておく。
	}
}

func (s *IntSet) Has(x int) bool { //IntSet型の構造体にメソッドを定義する、ポインタがレシーバになってるのは値を実際に書き換えるから、Hasは負じゃない値xをセットが含んでいるかを確認する
	word, bit := x/64, uint(x%64)
	return (word < len(s.words)) && (s.words[word]&(1<<bit) != 0)
}

func (s *IntSet) Add(x int) { // Add はセットに負ではない値xを追加します。
	word, bit := x/64, uint(x%64) //xの商と余剰をそれぞれwordとbitに入れる
	for word >= len(s.words) {    //最初はs.wordsは空。s.wordsのlenは0で、xが1だとwordは0なので条件を満たす
		s.words = append(s.words, 0) //0をリストに追加する。lenが1になる。なので１回でループを抜ける。65以上から2つ目の要素が追加される。1つの要素で64個まで整数を保管。
	}
	s.words[word] |= 1 << bit //保管した数値番目のbitを1にしておく。
}

func (s *IntSet) UnionWith(t *IntSet) { // UnionWith は、 s と t の和集合を s に設定します。
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string { // String は"{1 2 3}"の形式の文字列としてセットを返します。
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int { //要素を取り出す => セットの中からbitが1になってる数を数える
	result := 0
	for _, v := range s.words {
		result += bits.OnesCount64(v) //bitが1の数を数える
	}
	return result
}

func (s *IntSet) Remove(x int) { // セットから x を取り除きます
	word, bit := x/64, uint(x%64) //xの商と余剰をそれぞれwordとbitに入れる。
	if word < len(s.words) {      //引数で受け取った数字の商分要素数が存在しているかを確認
		s.words[word] &^= (1 << bit) //あれば、その要素の中から該当するbitを落とす
	}
}

func (s *IntSet) Clear() { // セットからすべての要素を取り除きます
	for i := range s.words {
		s.words[i] = s.words[i] & 1 << 63 //64bitの最上位のみを1にする
		s.words[i] &^= (1 << 63)          //最上位のbitを落として、s.wordsの数だけループさせる。
	}
}

func (s *IntSet) Copy() *IntSet { // セットのコピーを返します
	var copyed IntSet
	for i := range s.words {
		copyed.words = append(copyed.words, s.words[i])
	}
	return &copyed
}
