package main

import "fmt"

func test(x int) {
	a := 10 / x
	fmt.Println("a=", a)
}

func main() {
	//defer延迟调用 main函数结束前调用
	defer fmt.Println("aaaaaaaaaa") //如果一个语句有多个defer语句他会以先进后出的顺序执行
	defer fmt.Println("bbbbbbb")    //哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
	// defer  test(0)
	defer fmt.Println("cccccccc")

	a := 10
	b := 20
	defer func() {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}() //代表调用此匿名函数
	a = 111
	b = 222
	fmt.Printf("外部a=%d,b=%d\n", a, b)

	q := 10
	w := 20
	defer func(q, w int) {
		fmt.Printf("q=%d,w=%d\n", q, w)
	}(q, w) //输出结果为10，20   把参数传递过去 传参的过程会先执行 已经传参只是没有调用等价于（10 ，20）
	q = 111
	w = 222
	fmt.Printf("外部q=%d,w=%d\n", q, w)

}
