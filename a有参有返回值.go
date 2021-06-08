package main

import "fmt"

func fu(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return

}

func main() {
	max, min := fu(2, 17)
	fmt.Printf("max=%d,min=%d\n", max, min)

	a, _ := fu(23, 45) //通过匿名变量丢弃某个返回值
	fmt.Println("a=", a)

}
