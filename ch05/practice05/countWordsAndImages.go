package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	num_words, num_imgs, err := CountWordsAndImages("https://xkcd.com/")
	if err != nil {
		fmt.Errorf("Requested URL may some errors: %s", err)
	} else {
		fmt.Println("Given URL's HTML file contains...")
		fmt.Printf("Words : %d\n", num_words)
		fmt.Printf("Images : %d\n", num_imgs)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	//imgの処理
	if (n.Type == html.ElementNode) && (n.Data == "img") {
		images++
	}

	//wordsの処理
	input := bufio.NewScanner(strings.NewReader(n.Data))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling { //今見ているnodeに子nodeがnilになるまで、子nodeをvisitに渡して再帰する。そして隣接するnodeをcにセットして、tree構造を下にたどっていく。
		temp_words, temp_images := countWordsAndImages(c) //この受け皿がないとノードが変わるたびに上書きされて累計値が出なかった。
		words += temp_words
		images += temp_images
	}
	return words, images
}
