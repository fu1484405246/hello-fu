package main

import "fmt"

//函数递归
func fu(a int) {
	if a == 1 { //函数终止的条件
		fmt.Println("a=", a)
		return //终止函数调用
	}
	fu(a - 1)
	fmt.Println("a=", a)
}

func sua(i int) int {
	if i == 100 {
		return 100
	}
	return i + sua(i+1)
}

func main() {
	fu(3) //调用函数
	sum := 100
	sum = sua(1)
	fmt.Println("sum=", sum)

}
