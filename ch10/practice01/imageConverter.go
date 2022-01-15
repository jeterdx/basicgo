package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	imageFormat := flag.String("format", "png", "option to determine the image format to change")
	flag.Parse()

	switch *imageFormat {
	case "gif":
		fmt.Println("ccccc") //下の関数は呼ばれているのに何故かここに入らない
		if err := toGif(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "Gif conversion fails : %v\n", err)
			os.Exit(1)
		}
	case "jpeg":
		fmt.Println("bbbb") //下の関数は呼ばれているのに何故かここに入らない
		if err := toJpeg(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "Jpeg conversion fails : %v\n", err)
			os.Exit(1)
		}
	case "png":
		fmt.Println("aaaaa") //下の関数は呼ばれているのに何故かここに入らない
		if err := toPng(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "Png conversion fails : %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Input a appropreate format type of images; gif, jpeg or png.")
		os.Exit(0)
	}
}

func toGif(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input Format =", kind)
	return gif.Encode(out, img, &gif.Options{NumColors: 5})
}

func toJpeg(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input Format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 3})
}

func toPng(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input Format =", kind)
	return png.Encode(out, img) //pngはno option
}
