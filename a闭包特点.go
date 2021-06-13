package main

import "fmt"

//函数返回值是一个匿名函数，返回一个函数类型
func fu1() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数 f来调用闭包函数
	f := fu1()
	fmt.Println(f()) //1
	fmt.Println(f()) ////4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
	fmt.Println(f()) //25
	//他不关心这些捕获了的变量和常量是否已经超越作用域
	//所以只有闭包还在使用他，这些变量就还会存在

}

func fu() int { //函数被调用时 x才分配空间 初始化为0
	var x int //没有初始化值为0
	x++
	return x * x //函数调用完毕 x自动释放
}

// 传统局部变量特点
func main01() {
	fmt.Println(fu())
	fmt.Println(fu())
	fmt.Println(fu())
	fmt.Println(fu())
	fmt.Println(fu())

	a := 10
	str := "qwe"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "cainiao"
		//fmt.Printf("内部a=%d,str=%s\n",a,str)
	}() //（）代表直接调用
	//fmt.Printf("外部a=%d,str=%s\n",a,str)

}
