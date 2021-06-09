package main

import "fmt"

//函数调用流程先调用后返回 先进后出
//函数递归，函数调用自己本身

func funcc(c int) { //7开始执行这个函数
	fmt.Println("c=", c) //8 输出c=1
}

func funcb(b int) { //5，开始执行
	funcc(b - 1)         //6,调用funcc这个函数 c=b-1
	fmt.Println("b=", b) //9， 输出b=2

}

func funca(a int) { //3,开始执行funca这个函数
	funcb(a - 1)         //4，调用 funcb这个函数   b=a-1
	fmt.Println("a=", a) //10 , 输出a=3

}

func main() { //1，程序从main（）函数开始执行
	funca(3)               //2，调用函数  a=3
	fmt.Println("cainiao") //11 , 输出（）内容

	//输出结果为 c=1
	//         b=2
	//         a=3
	//         cainiao

}
