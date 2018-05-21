/*
go的函数是编程，函数是一等公民
go语言的函数的返回值的个数可以是多个值。
go里面是没有默认参数，重载什么的
但是go语言是有可变参数列表的
*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func operation(op func(int, int) int,a, b int) int {
	//打印出op的值
    p := reflect.ValueOf(op).Pointer() // 此时调用的p的指针内容是： 4761920
    opName := runtime.FuncForPC(p).Name()//获取到这个指针的ame, 此时调用的p的指针内容是： main.add
    fmt.Println("此时调用的p的指针内容是：", opName)
    return op(a, b)
}


func add(m, n int) int {
	return m + n
}

//定义返回值的默认名字，方便编译器自动生成
func div(a,b int)(q,r int){
	return a / b, a % b
}

//返回多个值的实际使用的方式就是一个返回值存应该接收的参数，另一个返回值接收返回的异常的参数即result和error
func operations1(a,b int, op string) (int, error)  {
	switch op {
	case "+":
		return a+b, nil
	case "-":
		return a-b, nil
	case "*":
		return a*b, nil
	case "/":
		a, _ = div(a, b)//这个地方不可以使用q接收参数的结果
		return a, nil
	default:
		return 0,fmt.Errorf("%s运算符有问题", op)
	}
}

//可变参数列表的操作与使用
func sum(args ...int) int {
	s := 0
	fmt.Println(args)//[1 2 3 4 5]
	//for i:=range args {//遍历args
	//	s += args[i]
	//}
	for i:=0; i<len(args);i++{//遍历args下标
	   s += args[i]
	}
	return s
}

func main() {
	v := operation(add, 5, 6)
	fmt.Println(v)

	//使用匿名函数，go的lambda就是比较极简
	v2 := operation(func(i,j int) int{
		return i+j
	}, 6, 7)// 此时调用的p的指针内容是： main.main.func1,因为我是一个匿名函数，所以在这里我们叫做func1
	fmt.Println(v2)//13

	q,r := div(13,2)
	fmt.Println(q,r)

	value, err := operations1(3, 5, "2")
    if err != nil{
    	fmt.Println(err)
	}else {
		fmt.Println(value)
	}

	su := sum(1,2,3,4,5)
    fmt.Println(su)//15
}
