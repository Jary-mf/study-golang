package main

import (
	"bufio"
	"fmt"
	"os"
)

func Word_main() {
	counts := make(map[string]int)
	//命令行要输入文件路径
	for _, arg := range os.Args[1:] {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Println("Open File Error")
			os.Exit(1)
		}
		input := bufio.NewScanner(file)
		input.Split(bufio.ScanWords)
		for input.Scan(){
			counts[input.Text()]++
		}
	}

	

}
