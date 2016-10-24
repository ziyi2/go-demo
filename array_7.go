package main

import "fmt"

func main() {
	/**
	一维数组
	 */
	//声明数组
	//Go语言数组声明需要指定元素类型及元素个数
	//var variable_name [SIZE] variable_type
	//例如 var balance [10] float32
	//初始化数组 var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//如果忽略[]中的数字不设置数组大小，Go语言会根据元素的个数来设置数组的大小：
	//var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//该实例与上面的实例是一样的，虽然没有设置数组的大小。

	//声明
	var n [10]int //n是一个长度为10的数组
	var i int

	for i = 0; i < 10; i++ {
		fmt.Printf("n[%d] = %d\n", i, n[i])		//初始化为0
	}

	/**
	多维数组
	 */
	//声明方式
	//var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	//var threedim [5][10][4]int

	var a [3][4]int //声明

	//二维数组初始化
	a = [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /*  第三行索引为 2 */
	}

	fmt.Printf("a[2][2] = %d\n", a[2][2])		//10

	var m,l int

	/* 输出数组元素 */
	for  m = 0; m < 3; m++ {
		for l = 0; l < 4; l++ {
			fmt.Printf("a[%d][%d] = %d\n", m,l, a[m][l] )
		}
	}


	/**
	向函数传递数组
	 */
	//1.形参设定数组大小
	//void myFunction(param [10]int)

	//2.形参未设定数组大小
	//void myFunction(param []int)

	var arr = []int {1,2,3,4,5}
	var avg float32

	avg = getAverage(arr, 5)

	fmt.Printf( "平均值为: %f ", avg );		//3.000000

}


func getAverage(arr []int, size int) float32 {
	var i,sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)	//强制类型转换?

	return avg
}



