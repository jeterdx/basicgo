package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値z を表している。
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0xcc - contrast*n, 0x11 - n, 0, 0x8a} //めちゃくちゃ適当に値を変えてみただけ。意味はない。
		}
	}
	return color.RGBA{0xaa, 0, 0xad, 0xfd} //めちゃくちゃ適当に値を変えてみただけ。意味はない。
}
