package main

import "fmt"


/**
全局变量
*/
var g int



func main() {
	/**
	局部变量
	在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
	 */
	//main函数中的局部变量

	//声明
	var a,b,c int

	//初始化
	a = 10
	b = 20
	c = a + b

	fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)


	/**
	全局变量
	 */
	g = a + b
	fmt.Printf ("g = %d\n", g)	//30

	//全局变量和局部变量名称可以相同, 但是局部变量会优先被考虑

	var g int = 10	//局部g
	fmt.Printf ("g = %d\n", g)	//10

	/**
	形式参数
	形式参数会作为函数的局部变量来使用
	 */
}
