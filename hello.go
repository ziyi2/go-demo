//第一行代码package main定义了包名。
//你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
//package main表示一个可独立执行的程序，每个Go应用程序都包含一个名为main的包
package main


//告诉Go编译器这个程序需要使用fmt包（的函数，或其他元素）
//fmt包实现了格式化IO（输入/输出）的函数
//是不是类似于Js中ES6语法的import 'path'
import (
	"fmt"
	"go/types"
)

//单个变量声明
//方式一: var v_name v_type 例如变量a,b 声明后若不赋值使用默认值
//方式二: var v_name = value 例如变量c,根据值自动判断变量类型
//方式三: 省略var v_name := value v_name不能被声明过,例如变量d,只能在函数体内出现

var a bool
var b string = "Hello, Variable"
var c = 10 	//和javascript有啥区别?

//多个变量声明,一般用于非全局变量
var e,f bool
var g,h = 1,'1'

//k,l := 1,'2'




// 这种因式分解关键字的写法一般用于声明全局变量
var (
	m string
	n int
)



//程序开始执行的函数
//main函数是每一个可执行程序所必须包含的
//一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数
//看起来是不是类似于C语言的int  main(void)
func main() {
	//这种声明方式只能在函数体内出现
	d := 10;
	i,j := 1,'2'


	fmt.Println("Hello, World!")
	fmt.Print("Hello, World!\n")		//和第一句功能一样
	fmt.Println("Hello, Go!")
	println(a,b,c,d)			//false "Hello, Variable" 10 10
	println(e,f)				//false false
	println(g,h)				//1 49
	println(i,j)				//1 50
	//println(k,l)				//non-declaration statement outside function body
	println(m)				//''
	println(n)				//0


	//值类型 int bool float string
	var intA = 11
	var intB int
	intB = intA
	println(intB)			//11
	println(&intA)			//0xc042031ec0 就是C语言...
	println(&intB)			//0xc042031eb8 地址不一样,说明开辟了新的内存块

	var boolA = true;
	var boolB = boolA;
	println(boolB)			//true
	println(&boolA)			//0xc042031eb7
	println(&boolB)			//0xc042031eb6 地址不一样,说明开辟了新的内存块

	//引用类型





}

