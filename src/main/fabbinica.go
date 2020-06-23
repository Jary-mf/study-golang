package main

import (
	"fmt"
	"os"
	"strconv"
)

func fabb(key int) int {

	if key == 1 || key == 0 {
		return 1
	}
	return fabb(key-1) + fabb(key-2)
}

func Fabb_main() {

	for _, key := range os.Args[1:] {
		fa, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(fabb(fa))
	}
}
