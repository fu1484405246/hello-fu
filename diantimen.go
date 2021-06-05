package main

import "fmt"

func main() {
	var men int
	fmt.Println("请按楼层:")
	fmt.Scan(&men)
	switch men { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是一楼")
	case 2:
		fmt.Println("按下的是二楼")
	case 3:
		fmt.Println("按下的是三楼")
	case 6, 7, 8: //case后面可以放条件
		fmt.Println("按下的是六六六楼")
	default: //其他情况
		fmt.Println("按下的是16楼")

	}
	a := 86
	switch {
	case a > 90:
		fmt.Println("优秀")
	case a > 80:
		fmt.Println("良好")
	case a > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其他")

	}
}
