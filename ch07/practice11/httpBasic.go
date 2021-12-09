package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

//dbに登録されている製品一覧を表示する
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

//価格と製品名をパラメータで受け取り、dbに追加する
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item] //存在していたらすでに登録済みであるメッセージを出力して関数を終了する
	if ok {
		fmt.Fprintf(w, "The item is already registered, please update the price information if you want: %q\n", item)
		return
	}
	textprice := req.URL.Query().Get("price")          //値段をパラメータで受け取って
	floatprice, _ := strconv.ParseFloat(textprice, 32) //floatに変換し
	db[item] = dollars(floatprice)                     //dollarsにキャストしてdbに追加
	fmt.Fprintf(w, "%s has been successfully added to the list at %s dollar.", item, db[item])
}

//パラメータで指定された商品の値段を表示する
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n Please create an item to the list at first.", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

//指定された製品の値段を更新する
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n Please create an item to the list at first.", item)
		return
	}
	textprice := req.URL.Query().Get("price")          //値段をパラメータで受け取って
	floatprice, _ := strconv.ParseFloat(textprice, 32) //floatに変換し
	db[item] = dollars(floatprice)                     //dollarsにキャストしてdbに追加
	fmt.Fprintf(w, "The price of %s has been successfully updated to %s dollar.", item, db[item])
}

//指定された製品の情報を削除する
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n Please create an item to the list at first.", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s has been successfully deleted from the list.", item)
}
