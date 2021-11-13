package main

import (
	"fmt"
)

// prereqs は情報科学の各講座をそれぞれの事前条件となる講座と対応付けします。
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": false},
	"calculus":   {"linear algebra": false},
	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},
	"data structures":       {"discrete math": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool) //関数型の変数を宣言

	visitAll = func(items map[string]bool) { //実際に関数を定義して代入
		for item := range items { //mapの中のmapに変更したので、通常にvalueをrangeする
			if !seen[item] { //seenはデフォルトでfalseなので、一度も参照されていなければtrueにする
				seen[item] = true //trueにして、
				visitAll(m[item]) //valueをkeyとして再帰呼び出し、keyとして存在せずseenがtrueになったら、以下でorderに順番に追加していく。
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = false //keysをmap型で使うためだけなのでなんでも良い。
	}

	visitAll(keys)
	return order
}
