package main

import "fmt"

func main() {
	a := 10
	str := "qwe"
	//匿名函数，没有函数名 函数定义
	fu := func() { //自动推导
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}
	fu()

	func() {
		fmt.Printf("a=%d,str=%v\n", a, str)
	}() //后面的（）代表调用此匿名函数

	//带参数的匿名函数
	func(a, b int) {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(12, 44)

	//匿名函数，有参有返回值
	q, w := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(20, 30)
	fmt.Printf("q=%d,w=%d\n", q, w)

}
