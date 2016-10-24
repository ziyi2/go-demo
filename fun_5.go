package main

import (
	"fmt"
	"math"
)


//结构体
type Circle struct {
	radius float64
}




func main() {

	/**
	普通函数
	 */


	var a = 13
	var b = 14
	var ret int

	ret = max(a,b)
	fmt.Printf("最大值是: %d\n", ret)	//14

	c,d := swap(a,b)
	fmt.Println(c,d)			//14 13


	/**
	函数参数
	形参就像定义在函数体内的局部变量。
	调用函数，可以通过两种方式来传递参数：
	值传递:值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递:引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	 */

	//值传递
	//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

	var e = 100
	var f = 200

	swapVaule(e,f)

	fmt.Printf("e = %d\n", e)	//e = 100
	fmt.Printf("f = %d\n", f)	//f = 200

	//引用传递,通过指针的形式
	swapPointer(&e,&f)			//传入的是地址

	fmt.Printf("e = %d\n", e)	//e = 200
	fmt.Printf("f = %d\n", f)	//f = 100


	/**
	函数作为值
	 */
	//声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	//使用函数
	fmt.Println(getSquareRoot(9))	//3

	/**
	函数闭包
	Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
	 */

	nextNumber := getSequence()

	fmt.Println(nextNumber())	//1
	fmt.Println(nextNumber())	//2

	//另外一个闭包
	newNextNumber := getSequence()
	fmt.Println(newNextNumber())	//1
	fmt.Println(newNextNumber())	//2


	/**
	Go语言函数方法
	Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
	func (variable_name variable_data_type) function_name() [return_type]{
   		 函数体
	}
	*/


	var g Circle
	g.radius = 10.00
	fmt.Println("Area of Circle(c) = " , g.getArea())	//314
}

/*
func function_name( [parameter list] ) [return_types] {
   函数体
}

//参数列表指定的是参数类型、顺序、及参数个数
//返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。

 */
func max(num1,num2 int) int {
	var result int

	if(num1 > num2) {
		result = num1
	} else {
		result = num2
	}

 	return result
}


//函数返回多个值
func swap(x,y int) (int, int) {
	return y,x
}


//交换值
func swapVaule(x,y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

//改变地址的值,两个参数是指针
func swapPointer(x *int, y *int) {
	var temp int
	temp = *x	//保持x地址上的值
	*x = *y		//将y赋给x (其实就改变x指针所指向的地址的值为y指针所指向的地址的值)
	*y = temp	//将temp值赋给y(这里当然也是改变y指针所指向的地址的值)
}



//闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}


//Circle类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}




