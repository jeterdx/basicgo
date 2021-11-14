package main

import (
	"fmt"
	"os"
)

//以下のパターンに対応できてない。
var prereqs = map[string]map[string]int{
	"algorithms":      {"data structures": 0},
	"data structures": {"discrete math": 0, "algorithms": 0},
	"discrete math":   {"intro to programming": 0, "algorithms": 0},
}

/*
	"algorithms": {"data structures": 0},
	"calculus":   {"linear algebra": 0},
	"compilers": {
		"data structures":       0,
		"formal languages":      0,
		"computer organization": 0,
	},
	"data structures":       {"discrete math": 0},
	"databases":             {"data structures": 0},
	"discrete math":         {"intro to programming": 0},
	"formal languages":      {"discrete math": 0},
	"networks":              {"operating systems": 0, "computer organization": 0},
	"operating systems":     {"data structures": 0, "computer organization": 0},
	"programming languages": {"data structures": 0, "computer organization": 0},
	"linear algebra":        {"calculus": 0},
}
*/

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string]map[string]int) []string {
	var order []string
	seen := make(map[string]int)
	var visitAll func(items map[string]int) //関数型の変数を宣言
	visitAll = func(items map[string]int) { //実際に関数を定義して代入
		for item := range items { //mapの中のmapに変更したので、通常にvalueをrangeする
			seen[item]++         //参照されるitemの値を1増やす
			if seen[item] == 1 { //一回目の参照であれば1なので、visitAllを再帰呼び出しさせる
				visitAll(m[item]) //valueをkeyとして再帰呼び出し、keyとして存在せずseenがtrueになったら、以下でorderに順番に追加していく。
				order = append(order, item)
				if seen[item] == 2 { //再帰呼び出し先でincrementされていた場合、循環としてプログラムを終了する。
					fmt.Println("トポロジカルソートが循環したためプログラムを終了しました。")
					os.Exit(1)
				}
			}
		}
	}

	keys := make(map[string]int)
	for key := range m {
		keys[key] = 0 //keysをmap型で使うためだけなのでなんでも良い。
	}

	visitAll(keys)
	return order
}
