package main

import "fmt"

type Account struct {
	author string
	password string
	age int
}



func main() {
	//Go语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
	//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

	//定义
	//type struct_variable_type struct {
	//   member definition;
	//   member definition;
	//   ...
	//   member definition;
	//}


	//声明
	//variable_name := structure_variable_type {value1, value2...valuen}


	var acc1 Account	//声明acc1为Account类型
	var acc2 Account

	acc1.author = "Go"
	acc1.password = "111"
	acc1.age = 9


	acc2.author = "Node"
	acc2.password = "111"
	acc2.age = 9


	fmt.Printf( "acc1 author : %s\n", acc1.author)
	fmt.Printf( "acc2 author : %s\n", acc2.author)

	/**
	结构体作为函数参数
	 */

	printAccount(acc1);
	printAccount(acc2);

	/**
	结构体指针
	 */
	//var struct_pointer *Books

	printAccountPointer(&acc1);
	printAccountPointer(&acc2);





}


func printAccount(acc Account) {
	fmt.Printf( "acc author : %s\n", acc.author)
}


//指针
func printAccountPointer(acc *Account) {
	fmt.Printf( "acc author : %s\n", acc.author)
}