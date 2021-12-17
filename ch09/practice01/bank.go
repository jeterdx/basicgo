// パッケージ bank は一つの口座を持つ並行的に安全な銀行を提供します。
package main

import "fmt"

func main() {
	for {
		go Deposit(10000)
		go fmt.Println(Balance())
		go Deposit(3000)
		go fmt.Println(Withdraw(2000))
		go fmt.Println(Balance())
		go fmt.Println(Withdraw(20000))
		go fmt.Println(Balance())
	}
}

var deposits = make(chan int)  // 入金額を送信する
var balances = make(chan int)  // 残高を受信する
var withdraws = make(chan int) //引き出し金額送信する
var result = make(chan bool)   //引き出しの結果を受信する

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-result
}

func teller() {
	var balance int // balance は teller ゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits: //deposits額をチャネルから受け取り、変数amountに代入する。宣言もここでしてる
			balance += amount
		case balances <- balance: //balanceは何もしてない。
		case amount := <-withdraws: //同じく、withdraw額をチャネルから受け取り、変数amountに代入してる。
			if balance >= amount { //もし、残高が引き出し金額以上であれば、成功で残額から引いておく。以下なら失敗。
				balance -= amount
				result <- true
			} else {
				result <- false
			}
		}
	}
}
func init() {
	go teller() // モニターゴルーチンを開始する、こいつはずっと稼働していて、selectの状態になった場合に、処理をかけていく
}
