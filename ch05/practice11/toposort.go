package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"}, // loop
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, ok := topoSort(prereqs)
	if !ok {
		log.Fatal("the graph has loops")
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

//Used a code from github.
func topoSort(m map[string][]string) (order []string, ok bool) { //boolでloopしているかしていないかを返す
	seen := make(map[string]bool)
	processing := make(map[string]bool) //処理中のセットをもう一つ用意

	var visitAll func(items []string) //関数型の変数を宣言
	ok = true                         //デフォルトはtrue

	visitAll = func(items []string) { //実際に関数を定義して代入、引数はstringのリスト
		for _, item := range items { //それぞれの前提科目を回す
			if processing[item] { //processing[item]がtrue、つまり、以下のvisitAllの再帰呼び出しが終わってから再度見られたら、それをloopとして判定する。初回は絶対にfalseなので入らない。
				ok = false
				return
			}
			if !seen[item] { //itemが見られていなかったらprocessing[item]をtrueにして処理中として判定、visitallを再帰呼び出しする。もし、再帰呼び出し先で次の同じitemが呼ばれるようであれば、上記の部分でloopの判定になる。
				processing[item] = true
				visitAll(m[item])

				order = append(order, item) //無事に経路が通り切れば、itemを追加し、processingのstatusもfalseに変更、seenもtrueにする。
				processing[item] = false
				seen[item] = true
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order, ok
}
