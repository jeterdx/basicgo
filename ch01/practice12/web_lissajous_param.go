package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black} //配列、スライス、マップ、構造体をコンポジットリテラルと呼ぶ。異なるデータ型をまとめて1つのデータ型にしたもののこと。この行ではスライス型が作成されていて、可変長な配列のことである。

const ( //プログラム内部で変わることのない定数を設定
	//whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) { //handlerの中にURLのクエリを受け付ける処理を追加
		keys, missing_message := r.URL.Query()["cycles"] //返り値の1つ目が入力された値、二つ目がエラー値
		if !missing_message || len(keys[0]) < 1 {
			log.Println("Url Param 'cycles' is missing")
			return
		}
		cycles, err := strconv.Atoi(keys[0]) //cyclesという変数にkeys[0]、つまり、urlで入力されたパラメータ値を整数型に変換して格納
		if err != nil {
			log.Println("Url Param 'cycles' is invalid") //エラーが起こった場合の対処
		}
		lissajous(w, cycles) //lissajous関数を実行、cyclesを使うために引数としてcyclesを取れるようにlissajousも書き換えている
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:2021", nil)) //webサーバの起動はここで行っている。
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
