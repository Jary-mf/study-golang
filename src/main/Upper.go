//词法域问题
package main

import "fmt"

func main() {
	x := "hello!" //此处定义了一个string类型的变量x，并赋值
	for i := 0; i < len(x); i++ {
		x := x[i] //此处又定义了一个作用域在for之内的变量x，
		//由于等式先执行右边，所以x[i]依然访问的是字符串x
		//z之后再将值赋给新定义的变量x
		if x != '!' { //寻找变量定义是从内到外，因此x为for中的x
			x := x + 'A' - 'a' //此处定义了一个作用域在if之内的变量x
			//之后同上，将for中的x+'A'-'a'赋值给新定义的x
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
			continue
		}
		fmt.Printf("%c", x)
	}

	fmt.Println()

	fmt.Println(x) //虽然实现了输出变大写，但原字符串没有变化

	f := 3e10
	i := int(f)
	fmt.Println(i)
	str:= "abcd"
	fmt.Println(len(str))
	fmt.Println("\xe4\xb8\x96")
}
