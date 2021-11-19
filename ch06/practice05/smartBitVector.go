package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

const intSize = 32 << (^uint(0) >> 63) //intsizeをいい感じにやってくれる賢い式

//以下64bit固定だったものをintSizeで代替
type IntSet struct {
	words []uint
}

func main() {
	var x IntSet
	var y IntSet
	x.Add(0)
	fmt.Println(intSize)
	y.AddAll(64, 66, 100, 122, 156, 1676, 1366)
	fmt.Println(intSize)

}

func (s *IntSet) Elems() []int {
	var elemList []int
	for index, sword := range s.words {
		for i := 0; i < intSize; i++ {
			if sword>>uint(i)&1 == 1 { //各要素のbitを1つずつ1かどうか判定
				elemList = append(elemList, intSize*index+i) //商+余剰で元に戻す
			}
		}
	}
	return elemList
}

func (s *IntSet) IntersectWith(t *IntSet) { //積集合、基本は和集合と同じ、それぞれのセットの要素数の違いを考慮に入れる
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) { //差集合、xからyを引く
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) { //対称差集合、二つの集合に被ってないもの
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) AddAll(xList ...int) { //可変個の引数を受け取ってセットに追加する。
	for _, v := range xList {
		word, bit := v/intSize, uint(v%intSize)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) Has(x int) bool { //IntSet型の構造体にメソッドを定義する、ポインタがレシーバになってるのは値を実際に書き換えるから、Hasは負じゃない値xをセットが含んでいるかを確認する
	word, bit := x/intSize, uint(x%intSize)
	return (word < len(s.words)) && (s.words[word]&(1<<bit) != 0)
}

func (s *IntSet) Add(x int) { // Add はセットに負ではない値xを追加します。
	word, bit := x/intSize, uint(x%intSize) //xの商と余剰をそれぞれwordとbitに入れる
	for word >= len(s.words) {              //最初はs.wordsは空。s.wordsのlenは0で、xが1だとwordは0なので条件を満たす
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
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int { //要素を取り出す => セットの中からbitが1になってる数を数える
	result := 0
	for _, v := range s.words {
		result += bits.OnesCount(v) //bitが1の数を数える
	}
	return result
}

func (s *IntSet) Remove(x int) { // セットから x を取り除きます
	word, bit := x/intSize, uint(x%intSize) //xの商と余剰をそれぞれwordとbitに入れる。
	if word < len(s.words) {                //引数で受け取った数字の商分要素数が存在しているかを確認
		s.words[word] &^= (1 << bit) //あれば、その要素の中から該当するbitを落とす
	}
}

func (s *IntSet) Clear() { // セットからすべての要素を取り除きます
	for i := range s.words {
		s.words[i] = s.words[i]&1<<intSize - 1 //64bitの最上位のみを1にする
		s.words[i] &^= (1<<intSize - 1)        //最上位のbitを落として、s.wordsの数だけループさせる。
	}
}

func (s *IntSet) Copy() *IntSet { // セットのコピーを返します
	var copyed IntSet
	for i := range s.words {
		copyed.words = append(copyed.words, s.words[i])
	}
	return &copyed
}
