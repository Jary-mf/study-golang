package main

import (
	. "bitmap"
	"fmt"
)

func Bitmap_main() {
	var x, y Bitmap
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Has(9))
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionBitMap(&y)
	fmt.Println(x.String())

	fmt.Println(x.Len())

	x.Remove(42)
	fmt.Println(x.String())

	y.Clear()
	fmt.Println(y.Has(9))

	y = x.Copy()
	fmt.Println(y.Has(1))

	ans := x.Elems()
	fmt.Println(ans)

}
