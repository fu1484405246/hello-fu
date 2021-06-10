package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型 通过type给一个函数类型起名
type futype func(int, int) int //futype：函数类型起名

func main() {
	var result int
	//result = add(2,5)
	//fmt.Println("result=",result)
	//result =minus(3,2)
	//fmt.Println("result=",result)
	var cainiao futype //声明一个函数类型的变量 变量名叫cainiao
	cainiao = add
	result = cainiao(10, 20) //等价于add（10，20）
	fmt.Println("result=", result)

}
