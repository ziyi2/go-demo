package main

import "fmt"


func main() {
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



