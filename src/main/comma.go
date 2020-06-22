package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma1(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]

}

func comma2(ss string) string {

	var flag string
	if ss[0] == '+' || ss[0] == '-' {
		flag = flag + string(ss[0])
		ss = ss[1:]
	}

	dot := strings.LastIndex(ss, ".")
	var s, ans string
	if dot == -1 {
		s = ss
		ans = s[len(s)-3 : len(s)]
	} else {
		s = ss[:dot+1]
		ans = s[len(s)-3:len(s)] + ss[dot+1:]
	}
	var buf bytes.Buffer
	buf.WriteString(flag)
	n := len(s)
	if n <= 3 {
		return s
	}
	var i int
	for i = n - 6; i > 0; i -= 3 {
		sub := s[i : i+3]
		ans = sub + "," + ans
	}
	if i <= 0 {
		ans = s[:i+3] + "," + ans
	}
	buf.WriteString(ans)
	return buf.String()
}

func Comma_main() {
	for _, arg := range os.Args[1:] {
		// str1 := comma1(arg)
		// fmt.Println(str1)
		str2 := comma2(arg)
		fmt.Println(str2)
	}
}
