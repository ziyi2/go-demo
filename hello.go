//第一行代码package main定义了包名。
//你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
//package main表示一个可独立执行的程序，每个Go应用程序都包含一个名为main的包
package main


//告诉Go编译器这个程序需要使用fmt包（的函数，或其他元素）
//fmt包实现了格式化IO（输入/输出）的函数
//是不是类似于Js中ES6语法的import 'path'
import "fmt"




//程序开始执行的函数
//main函数是每一个可执行程序所必须包含的
//一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数
//看起来是不是类似于C语言的int  main(void)

func main() {
	fmt.Println("Hello, World!");
	fmt.Print("Hello, World!\n");	//和第一句功能一样
	fmt.Println("Hello, Go!");
}

