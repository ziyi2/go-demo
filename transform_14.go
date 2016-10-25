package main

import "fmt"

func main() {
	//type_name(expression)
	//type_name 为类型，expression 为表达式。

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)



}
