package main

import (
	"html/template"
	"log"
	"net/http"
)

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("templates/svg.html"))

	str := "svg output"

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "svg.html", str); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", htmlHandler0)

	// サーバーを起動
	log.Fatal(http.ListenAndServe(":8989", nil))
}
