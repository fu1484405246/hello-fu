package main

import "fmt"
import "time"

func main() {
	i := 2
	for {
		i++
		time.Sleep(time.Second) // 延迟一秒输出
		if i == 5 {
			break //跳出循环，如果嵌套多个循环，跳出最近的那个循环（{   {  {break}  }   }）  输出结果： 3 4 结束本次循环
			//continue //(跳过本次循环)，下一次继续 输出结果：3 4 6 7 ~~~~~~一直循环
		} //continue只能用在循环
		fmt.Println("i=", i)

	}

	a := 0
	for i := 1; i <= 100; i++ { //初始条件； 判断条件； 条件变化;
		a = a + 1
	}

	goto End //goto是关键字  End是用户起的名字他叫标签

	fmt.Println("a=", a)

End:
	fmt.Println("菜鸟")

	//range 迭代打印每个元素，默认返回2个值:一个是元素的位置一个是元素本身
	str := "qwer"
	for i, data := range str { //data 返回元素（qwer）
		fmt.Printf("str[%d]=%c\n", i, data)
	}

}
