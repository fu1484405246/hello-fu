package main

import "fmt"

//有参无返回值函数的定义，普通参数列表
func cainiao(a int, b float64) { //固定参数
	//定义函数时，在函数名后面()定义的参数叫形参
	fmt.Printf("a=%d,b=%f"+
		"\n", a, b)

}

func fu(b ...int) { //这样传递的实参可以是0或者多个
	//...int类似这种类型实质上是...type不定参数类型
	fmt.Println("len(b)=", len(b)) //获取用户传递参数的个数
	//len: 内建函数
}

func main() {
	cainiao(22, 23.9) //有参无返回值调用： 函数名（所需参数）
	//调用函数传递的参数叫实参
	//参数传递，只能由实参传递给形参

	fu(233, 32, 32)

}
