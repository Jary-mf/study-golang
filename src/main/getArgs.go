// main.go
package main

import (
	"fmt"
	"os"
)

func main_main() {
	s:="echo"
	for _, arg := range os.Args[1:] {
		s += " " + arg
	}
    fmt.Println(s)
    fmt.Println(os.Args[0])
}
