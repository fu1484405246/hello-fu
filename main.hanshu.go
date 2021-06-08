package main

import "fmt"

func cainiao(a ...int) {
	for _, data := range a {
		fmt.Println("data=", data)
	}

}
func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data", data)
	}
}

func caigo(b ...int) {
	//把全部元素传递给另外一个函数（caogo传递给菜鸟）
	cainiao(b...)

	//当只想传递1，2时
	myfunc(b[:2]...) //输出结果为 1，2
	//myfunc(b[2:]...)//输出结果为 3，4

}
func main() {

	caigo(1, 2, 3, 4)

}
