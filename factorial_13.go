package main

import "fmt"

func main() {
	//func recursion() {
	//	recursion() /* 函数调用自身 */
	//}
	//
	//func main() {
	//	recursion()
	//}

	//Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
	//递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))



}


func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}
