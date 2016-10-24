package main

import "fmt"


func main() {
	var a = 21	//这里看起来是不是和JS一模一样
	var b = 10
	var c int

	c = a + b;
	fmt.Printf("a + b = %d\n", c)	//31

	a++
	fmt.Printf("a++  = %d\n", a)	//22

	var d = "string"
	var e = "string"

	if(d == e) {
		fmt.Printf("d = e \n")	//d == e
	}


	var f = 2			//0000 0010
	var g = 4			//0000 0100
	var h = f | g			//0000 0110
	fmt.Printf("h = %d\n",h)	//h = 6

	//* 指针运算符
	var i = 4
	var p *int

	p = &i	//指针,和C语言一模一样
	fmt.Printf("p的变量类型为 %T\n", p)	//*int
	fmt.Printf("*p = %d\n", *p)		//4


}
