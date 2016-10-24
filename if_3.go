package main

import "fmt"

func main() {
	//if

	var a = 10;
	if a < 20 {
		fmt.Printf("a = %d\n", a);
	}

	//switch
	//switch语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
	//switch语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break

	var grade string
	var marks = 90

	switch marks {
		case 90: grade = "A"	//不需要break
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}

	fmt.Printf("grade = %s\n", grade);	//grade = A

	switch  {
		case grade == "A": fmt.Printf("A\n")	//A
		default: fmt.Printf("Other\n")
	}

	//Type Switch

	var x interface{}
	switch i:= x.(type) {
		case nil:
			fmt.Printf("x的类型是: %T", i)		//x的类型是: <nil>
		case int:
			fmt.Printf("x的类型是int型")
		default:
			fmt.Printf("未知型")
	}


	//select 类似于通信的switch语句
	//select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
	//select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")	//no communication
	}



}
