package main

import "fmt"

const MAX int = 3


func main() {
	/**
	指针
	 */
	//变量的地址
	var a = 10
	fmt.Printf("&a = %x\n", &a)	//&a = c0420082a0

	//指针
	//一个指针变量可以指向任何一个值的内存地址, 它指向那个值的内存地址
	//类似于变量和常量，在使用指针前你需要声明指针
	//var var_name *var-type

	//var ip *int
	//var fp *float32

	//定义指针变量
	//为指针变量赋值
	//访问指针变量中指向地址的值

	var ip *int

	ip = &a


	// Print 将参数列表 a 中的各个参数转换为字符串并写入到标准输出中。
	// 非字符串参数之间会添加空格，返回写入的字节数。
	//func Print(a ...interface{}) (n int, err error)

	// Println 功能类似 Print，只不过最后会添加一个换行符。
	// 所有参数之间会添加空格，返回写入的字节数。
	//func Println(a ...interface{}) (n int, err error)

	// Printf 将参数列表 a 填写到格式字符串 format 的占位符中。
	// 填写后的结果写入到标准输出中，返回写入的字节数。
	//func Printf(format string, a ...interface{}) (n int, err error)


	//fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	//fmt.Println("a", "b", 1, 2, 3, "c", "d")
	//fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	// ab1 2 3cd
	// a b 1 2 3 c d
	// ab 1 2 3 cd
	fmt.Printf("ip变量存储的地址是: %x\n", ip)	//和a变量的地址一样
	fmt.Printf("ip变量存储的地址所指向的值是: %d\n", *ip)		//10


	*ip = 11
	fmt.Printf("a = %d\n", a)		//11 a的值变化了,应为a的地址所指向的内存的值被修改了

	/**
	空指针
	 */
	var ptr *int
	fmt.Printf("ptr的值为: %x\n", ptr)	//0

	if(ptr == nil) {
		fmt.Printf("ptr是空指针\n")	//nil代表空指针
	}

	/**
	指针数组
	 */
	//如果需要保存数组
	//var ptr [MAX]*int 声明整形指针数组

	var i int
	var ptr1 [MAX]*int
	arr := []int{10,100,1000}

	for i = 0; i < MAX; i++ {
		ptr1[i] = &arr[i]	//整型地址赋值给指针数组
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr1[i])
	}
}



