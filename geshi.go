package main

import "fmt"

func main() {
	a := 3
	b := 2.4
	c := "qwer"
	d := 'a'
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d) //%T操作变量所属类型
	fmt.Printf("a=%d,b=%f,c=%s,d=%c\n", a, b, c, d)
	//%d整形格式打印
	//%f浮点个数打印
	//%s字符串格式打印
	//%c字节个数打印
	//%v自动匹配格式输出
	fmt.Printf("a=%v,b=%v,c=%v,d=%v\n", a, b, c, d)

	//输入的使用
	var x float64
	fmt.Println("输入x：")

	fmt.Scan(&x) /////阻塞等待用户的输入
	fmt.Println("x=", x)
}
