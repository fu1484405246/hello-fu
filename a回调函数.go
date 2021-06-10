package main

import "fmt"

type futype func(int, int) int

//回调函数：函数有一个参数是函数类型，这个函数就是回调函数
//多态 多种形态，调用同一个接口不同的表现，同一个接口可以实现不同表现
//先有想法在来实现功能
func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func Calc(a, b int, cainiao futype) (result int) {
	fmt.Println("Calc")
	result = cainiao(a, b)
	return

}

func main() {
	a := Calc(2, 6, add) //mul
	fmt.Println("a=", a)

}
