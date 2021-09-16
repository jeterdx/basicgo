package main

import (
	"fmt"
	"os"
	"strconv"

	"practice02/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		if arg == "" {
			//stdinを処理するコードを書く
		}
		tempcalc(t)
		lengthcalc(t)
		weightcalc(t)
	}
}

func tempcalc(t float64) {
	f := unitconv.Fahrenheit(t)
	c := unitconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToC(f), c, unitconv.CToF(c))
}

func lengthcalc(t float64) {
	m := unitconv.Metre(t)
	f := unitconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FtoM(f), m, unitconv.MToF(m))
}

func weightcalc(t float64) {
	k := unitconv.Kgm(t)
	l := unitconv.Lb(t)
	fmt.Printf("%s = %s, %s = %s\n", k, unitconv.KtoL(k), l, unitconv.LToK(l))
}
