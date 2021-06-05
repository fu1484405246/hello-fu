package main

import "fmt"

func main(){
	fmt.Println("4<5 结果为:",4<5)  //真
	fmt.Println("!(4<5) 结果为:",!(4<5)) //假
	//非：如果a为真，则a为假 如果a为假则a为真

	fmt.Println("4<5&&2<8 结果为" , 4<5&&2<8)
	fmt.Println("true&&false结果为:",true&&false)
	//&&与，并且，左右结果都为真结果才为真


	fmt.Println("true||false结果为：",true||false)
	fmt.Println("false||false结果为：",false||false)
	//或 || 左右俩边都为假结果才为假


	}





}
