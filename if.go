package main

import "fmt"

func main() {
	a := "符升"
	//if和{就是条件，条件通常都是关系运算符  括号和if是在同一行
	if a == "符升" {
		fmt.Println("菜鸟")
	}
	//if支持1个初始化语句，初始化语句和判断条件以分号分隔
	if a := 2; a == 2 { //条件为真执行括号里的语句
		fmt.Println("a==2")
	}

	if a := 25; a == 29 {
		fmt.Println("a==29")
	} else if a > 29 {
		fmt.Println("a>29")
	} else if a < 29 {
		fmt.Println("a<29")
	}

}
