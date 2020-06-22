package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup1_main() {
	counts := make(map[string]int)
	//命令行要输入文件路径
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println("ReadFile error!")
		}
		for _, word := range strings.Split(string(data), " ") {
			counts[word]++
			if counts[word] > 1 {
				fmt.Printf("%s exists in %s\n", word, arg)
			}
		}
	}
	for k, v := range counts {
		fmt.Printf("%s : %d\n", k, v)
	}
}
