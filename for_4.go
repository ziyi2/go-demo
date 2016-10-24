package main

import "fmt"

func main() {

	//永真循环
	//for true {
	//	fmt.Printf("这是一个无限循环.\n");
	//}

	//for循环
	//for init condition post
	//init 一般为赋值表达式，给控制变量赋初值；
	//condition 关系表达式或逻辑表达式，循环控制条件；
	//post  一般为赋值表达式，给控制变量增量或减量。
	for a := 0; a < 5; a++ {
		fmt.Printf("a =  %d\n", a)
	}

	//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	//for key, value := range oldMap {
	//	newMap[key] = value
	//}

	numbers := [6]int{1,2,3,4}

	for i,x := range numbers {
		fmt.Printf("x[%d] =  %d\n", i, x)
	}	//x[0] = 1 x[1] = 2 x[2] = 3 x[3] = 4 x[4] = 0 x[5] = 0


	//break在case中仍然也是跳出语句的作用



}