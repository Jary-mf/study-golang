package main

import "fmt"

func counters(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarers(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printers(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Channels_main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counters(naturals)
	go squarers(squares, naturals)
	printers(squares)
}
