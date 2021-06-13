package main

import "fmt"

var a byte // 全局变量

//不同作用域允许定义同变量名
//使用变量的原则就近原则

func main() {
	var a int //局部变量
	fmt.Printf("%T\n", a)
	{
		var a string
		fmt.Printf("%T\n", a)
	}

	fu()
}
func fu() {
	fmt.Printf("%T\n", a) //
}
