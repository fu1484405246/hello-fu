package main

import "fmt"

func main1() {
	//局部变量：定义在{}里面的变量，只能在{}里面有效
	//局部变量特点 执行到定义变量那句话，才开始分配空间，离开作用域会自动释放
	//作用域 变量其作用的范围
	a := 100
	fmt.Println("a=", a)

	//   a == 20 只能在{}里有效
}

var b int //定义在函数外的变量 叫全局变量
//全局变量在在如何地方都能使用
func main() {
	b = 12
	fmt.Println("b=", b)
	fu()

}
func fu() {
	fmt.Println("b=", b) //全局变量在在如何地方都能使用
}
