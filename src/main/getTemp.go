package main

import (
	"fmt"
	"os"
	"strconv"
	. "tempconv"
)

func Temp_main() {
	fmt.Println(CToF(BoilingC))
	fmt.Println(CToF(AbsoluteZeroC))
	fmt.Println(KToF(AbsoluteZeroK))

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s , %s = %s \n", f, FToC(f), c, CToF(c))

	}
}
