package main

import "fmt"

func main() {
	var a int
	var q byte
	q = 'b'
	a = int(q) //类型转换
	fmt.Println("a=", a)

	//类型别名
	type fusheng int64
	var w fusheng
	fmt.Printf("w type is %T\n", w)

	type (
		ge   int
		niao byte
	)
	var r ge = 11
	var e niao = 'a'
	fmt.Printf("r=%d,e=%c\n", r, e)

}
