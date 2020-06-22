package main

import (
	"crypto/sha256"
	"fmt"
)

func Diff(c1 [32]byte ,c2 [32]byte) int {
	
	var count int
	for i:=0;i<32;i++{
		k1,k2:=c1[i],c2[i]
		for j:=0;j<8;j++{
			if k1&1!=k2&1{
				count++
			}
			k1>>=1
			k2>>=1
		}
	}
	return count

}

func Diff_main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%x\n",[]byte("x"))
	fmt.Println(Diff(c1,c2))
}
