package main

import "fmt"

//无参有返回值：当只有一个返回值时：
func myfunc() int {
	return 20 //有返回值的函数需要通过return返回（中断函数）
}

func fu() (cainiao int) { //给返回值起一个变量名，go推荐写法
	cainiao = 12
	return
}

func cainiao() (int, int, int) {

	return 22, 33, 44

}

func main() {
	var a int
	a = myfunc()
	fmt.Println("a=", a)
	b := fu()
	fmt.Println("b=", b)
	q, w, e := cainiao()
	fmt.Printf("q=%d,w=%d,e=%d\n", q, w, e)
}
