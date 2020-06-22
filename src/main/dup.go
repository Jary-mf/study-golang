package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	//bufio.NewScanner()无限输入，按CTRL+Z结束输入
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup_main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	//无文件名参数，则进入命令行输入，否则打开文件
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println("Stdin error!")
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for k, v := range counts {
		fmt.Printf("%s : %d\n", k, v)
	}
}
