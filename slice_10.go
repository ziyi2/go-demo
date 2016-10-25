package main

import "fmt"

func main() {
	//Go语言切片是对数组的抽象
	//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

	//创建切片
	//var identifier []type
	//var slice1 []type = make([]type, len)
	//slice1 := make([]type, len) 简写
	//make([]T, length, capacity)	指定容量,其中capacity为可选参数,len是数组的长度并且也是切片的初始长度。

	//切片初始化
	//s :=[] int {1,2,3 } 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

	//s := arr[:]
	//s := arr[startIndex:endIndex]
	//初始化切片s,是数组arr的引用

	//s := arr[startIndex:] 缺省endIndex时将表示一直到arr的最后一个元素
	//s := arr[:endIndex]  缺省startIndex时将表示从arr的第一个元素开始

	//s1 := s[startIndex:endIndex]  通过切片s初始化切片s1

	//s :=make([]int,len,cap) 通过内置函数make()初始化切片s,[]int标识为其元素类型为int的切片


	//len() 和 cap() 函数
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

	var numbers = make([]int,3,5)
	printSlice(numbers);	//3 5 [0 0 0]

	/*
	nil切片
	 */
	var slice1 []int;

	printSlice(slice1);		//len=0 cap=0 slice=[]

	if(slice1 == nil) {
		fmt.Printf("切片是空的")
	}

	/*
	切片截取
	 */
	slice2 := []int{1,2,3,4,5,6,7,8,9}
	printSlice(slice2);		//len=9 cap=9 slice=[1 2 3 4 5 6 7 8 9]

	//索引1到4(不包含4,这和javascript很像)
	fmt.Println("slice2[1:4] ==", slice2[1:4])	//slice2[1:4] == [2 3 4]

	//默认下限为0
	fmt.Println("slice2[:4] ==", slice2[:4])	//slice2[1:4] == [1 2 3 4]

	//默认上限为 len(s)
	fmt.Println("slice2[4:] ==", slice2[4:])	//slice2[1:4] == [5 6 7 8 9]

	//子切片
	slice3 := slice2[:2]
	printSlice(slice3)	//len=2 cap=9 slice=[1 2]

	/**
	append()和copy()
	 */
	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	//下面的代码描述了从拷贝切片的copy方法和向切片追加新元素的append方法

	var slice4 []int

	printSlice(slice4);	//len=0 cap=0 slice=[]

	//追加空切片
	slice4 = append(slice4,0)
	printSlice(slice4);	//len=1 cap=1 slice=[0]

	//像非空切片添加一个元素
	slice4 = append(slice4,1)
	printSlice(slice4);	//len=2 cap=2 slice=[0 1]

	//同时添加多个元素
	slice4 = append(slice4, 2,3,4)
	printSlice(slice4)	//len=5 cap=6 slice=[0 1 2 3 4]

	//创建切片是之前切片的两倍容量
	slice5 := make([]int, len(slice4), (cap(slice4))*2)
	copy(slice5,slice4)
	printSlice(slice5)	//len=5 cap=12 slice=[0 1 2 3 4]

}


func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

