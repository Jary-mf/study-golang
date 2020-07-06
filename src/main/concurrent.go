package main

import (
	"fmt"
	"time"
)

var fibnum []int

func spinner(dalay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(dalay)
		}
	}
}

func fib(x int) {
	if x < 2 {
		return
	}
	for i := 2; i <= x; i++ {
		fibnum[i] = fibnum[i-1] + fibnum[i-2]
	}
}

func initfib(n int) {
	for i := 0; i <= n; i++ {
		fibnum = append(fibnum, 0)
	}
	fibnum[1] = 1
}

func Concurrent_main() {
	go spinner(100 * time.Millisecond)
	const n = 560
	initfib(n)
	fib(n)
	time.Sleep(5 * time.Second)
	fmt.Printf("\r Fibonacci(%d) is %d\n", n, fibnum[n])

}
