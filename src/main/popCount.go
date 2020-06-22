package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//循环查表
func PopCount1(x uint64, start time.Time, ch chan<- string) {

	var ans int
	for i := 0; i < 8; i++ {
		j := i * 8
		ans += int(pc[byte(x>>byte(j))])
	}
	ch <- fmt.Sprintf("Popcount1 ans : %d\t time : %f", ans, time.Since(start).Seconds())
}

//普通查表
func PopCount2(x uint64, start time.Time, ch chan<- string) {
	ans := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	_ = ans
	ch <- fmt.Sprintf("Popcount2 ans : %d\t time : %f", ans, time.Since(start).Seconds())
}

//按位与1
func PopCount3(x uint64, start time.Time, ch chan<- string) {

	var ans int
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ans++
			x = x >> 1
		}
	}
	ch <- fmt.Sprintf("Popcount3 ans : %d\t time : %f", ans, time.Since(start).Seconds())
}

//遇1清零 （超过64位失效）
func PopCount4(x uint64, start time.Time, ch chan<- string) {

	var ans int
	for x != 0 {
		x = x & (x - 1)
		ans++
	}
	ch <- fmt.Sprintf("Popcount4 ans : %d\t time : %f", ans, time.Since(start).Seconds())
}

func Pop_main() {
	ch1 := make(chan string)
	 ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	for _, arg := range os.Args[1:] {
		start := time.Now()
		x, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Need an Integer")
			continue
		}
		go PopCount1(uint64(x), start, ch1)
		fmt.Println(<-ch1)
		go PopCount2(uint64(x), start, ch2)
		fmt.Println(<-ch2)
		go PopCount3(uint64(x), start, ch3)
		fmt.Println(<-ch3)
		go PopCount4(uint64(x), start, ch4)
		fmt.Println(<-ch4)

		fmt.Printf("Main : %f", time.Since(start).Seconds())
	}
}
