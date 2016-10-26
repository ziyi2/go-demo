package main

import (
	"fmt"
	"errors"
)


//实例一 自定义Error
//error类型是一个接口类型，这是它的定义





type error interface {
	Error() string
}


//我们可以在编码中通过实现 error 接口类型来生成错误信息。
//函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}


	return f,errors.New("");	//注意这离需要返回两个值
}



//实例二 内置Error

//Go语言通过内置的错误接口提供了非常简单的错误处理机制。



//定义一个DivideError结构
type DivideError struct {
	dividee int
	divider int
}


//实现error接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat,de.dividee)
}

//定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}



func main() {

	//实例一
	result, err:= Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}




	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}






}


