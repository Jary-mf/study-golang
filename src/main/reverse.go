package main

import (
	"fmt"
)

func reverse(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s *[32]int) {
	tem := s[len(s)-1]
	for i := len(s) - 1; i > 0; i-- {
		s[i] = s[i-1]
	}
	s[0] = tem
}

func distinct(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}
	fmt.Printf("%x\n", s)
}

func Reverse_main() {
	var a [32]int
	for i := 0; i < 32; i++ {
		a[i] = i
	}
	fmt.Println(a[0])
	p := &a
	reverse(p)
	fmt.Println(a[0])
	rotate(p)
	rotate(p)
	fmt.Println(a[0])
	b := [6]int{1, 1, 2, 2, 3, 3}
	distinct(b[:])
	fmt.Println(fmt.Sprintf("%q\n",b[:]))
}
