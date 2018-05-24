/*
别名实现了go语言的多态
*/
package main

import "fmt"

func add(a, b int) int {
	return a+b
}

func min(a, b int) int {
	return a-b
}

func main() {
	// 类型别名的关键字就是一个type
	type bigint int64

	var a bigint
	// a type is main.bigint
	fmt.Printf("a type is %T\n", a)

	// 用别名定义一个匿名的函数，myFunc此时是一个函数类型，代表一个传入两个int参数，返回一个int值的函数类型的统称
	type myFunc func(int, int) int
	// 定义一个与此函数相同的形参与返回值的函数名
	var myFunc1 myFunc
	myFunc1 = add
	va := myFunc1(3 ,5)
	fmt.Println(va)//8

	myFunc1 = min
	va = myFunc1(3 ,5)
	fmt.Println(va)//-2
}
