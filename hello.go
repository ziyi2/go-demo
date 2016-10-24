//第一行代码package main定义了包名。
//你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
//package main表示一个可独立执行的程序，每个Go应用程序都包含一个名为main的包
package main


//告诉Go编译器这个程序需要使用fmt包（的函数，或其他元素）
//fmt包实现了格式化IO（输入/输出）的函数
//是不是类似于Js中ES6语法的import 'path'
import (
	"fmt"
	//"go/types"
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
	d := 10
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

	var boolA = true
	var boolB = boolA
	println(boolB)			//true
	println(&boolA)			//0xc042031eb7
	println(&boolB)			//0xc042031eb6 地址不一样,说明开辟了新的内存块

	//引用类型
	var arrA = [10]int{0,1,2,3,4,5,6,7,8,9}
	var arrB [10]int
	println(arrA[1])			//1

	arrB = arrA
	println(arrB[1])			//1

        arrA[1] = 1111
	println(arrA[1])			//1111
	println(arrB[1])			//1

	//常量,和ES6语法很像
	const ERROR string = "ERROR"
	const SUCCESS = "SUCCESS" 		//隐藏类型定义string
	const ERR,SUCC = "ERR","SUCC"		//多个变量同时声明

	println(ERROR);
	println(SUCCESS);
	println(ERR);
	println(SUCC);

	//枚举常量
	const (
		ZERO = 0
		FIRST = 1
		SECOND = 2
	)

	println(ZERO);
	println(FIRST);
	println(SECOND);


	//枚举常量,通过内置函数赋值(必须是内置的)

	const (
		NAME = "ZIYI2"
		NAME_LENGTH = len(NAME)	//len cap unsafe.Sizeof
	)

	println(NAME_LENGTH);	//5


	//iota 可以被编译器修改的常量 被用于枚举
	const (				//在每一个const关键字出现时被重置为0
		IOTA_0 = iota		//没出现一次iota,其所代表的数字会自动加 +1
		IOTA_1 = iota
		IOTA_2 = iota
	)

	println(IOTA_0);		//0
	println(IOTA_1);		//1
	println(IOTA_2);		//2


	//简写形式
	const (
		IOTA_00 = iota
		IOTA_11
		IOTA_22
	)

	println(IOTA_00);		//0
	println(IOTA_11);		//1
	println(IOTA_22);		//2

	//用法
	const (
		A = iota	//0
		B		//1
		C		//2
		D = "CONST_D"	//独立值 iota += 1    	iota = 3
		E		//iota += 1           	iota = 4
		F = 100		//iota += 1		iota = 5
		G		//iota += 1		iota = 6
		H = iota        //7 恢复计数
		I 		//8
	)

	fmt.Println(A,B,C,D,E,F,G,H,I)	//0 1 2 CONST_D CONST_D 100 100 7 8

	const (
		J = 1 << iota
		K = 3 << iota
		L
		M
	)

	fmt.Println("J = " , J); //1    (1<<0)	后面的数值是iota值
	fmt.Println("K = " , K); //6	(3<<1)  意思是数值3左移1位
	fmt.Println("L = " , L); //12	(3<<2)
	fmt.Println("M = " , M); //24 	(3<<3)
}

