package main

import "fmt"

var slice []func()

func trap()  {
	sli := []int{1, 2, 3, 4, 5}
	for _, v := range sli {
		fmt.Println(&v)
		slice = append(slice, func() {
			fmt.Println(v * v) // 直接打印结果
		})
	}

	for _, val := range slice {
		val()
	}
	fmt.Println(slice)
}

func noTrap (){
	sli := []int{1, 2, 3, 4, 5}
	for _, v := range sli {
		v:=v
		fmt.Println(&v)
		slice = append(slice, func() {
			fmt.Println(v * v) // 直接打印结果
		})
	}

	for _, val := range slice {
		val()
	}
	fmt.Println(slice)
}

func Trap_main() {
	//trap()
	noTrap()
}

// 输出 25 25 25 25 25
