/*
tips:1.在go语言中，定义的变量一定要使用到，否则就会报错（飘红）
原因是可以是我们在项目中不必定义很多无用的变量，会自动提示未使用到的垃圾变量
2.在go语言中可以省略赋值表达式左边变量的类型
3.声明变量时，我们可以省去var的关键字,让编译器自动决定类型
4.再定义局部变量时，我们一般都会使用:=的赋值方式
5.当在函数外部定义变量的时候，即定义的是包内部的变量，那么我们是不可以使用:=的方式赋值的
此时，我们就必须使用var 变量= 值的方式进行赋值
*/
package main

import (
	"fmt"
)

//初始声明的变量在未赋值的情况下，是有初始值的
func variable1(){
	var i int
	var s string
	// var 括号里声明可以自动类型识别
	var(
		a = 3
		b = 4.1
		c = '丁'
	)
	//fmt.Print(i, s)
	//格式化输出方式
	fmt.Printf("%d %q\n", i, s)//0 ""
	fmt.Printf("%d %f %c\n", a, b, c)//3 4.100000 丁
}

//为变量赋初始值
func variable2(){
	var i int = 3
	var s string = "Jefflikie"
	fmt.Printf("%d %q\n", i, s)// 3 "Jefflikie"
}

//为多个变量赋初始值
func variable3(){
	var i, j  int = 3, 8
	var s string = "Jefflikie"
	fmt.Printf("%d %d %q\n", i, j, s)// 3 8 "Jefflikie"
}

//声明多个变量
func variable4(){
	var i, j  int
	var s string = "Jefflikie"
	fmt.Printf("%d %d %q\n", i, j, s)// 0 0 "Jefflikie"
}

//赋值时可不必加上变量的具体类型，go会自行判断
func variable5(){
	var i, j  = 3, 9
	var s = "Jefflikie"
	fmt.Println(i, j, s)// 3 9 "Jefflikie"
}

// 可以达到类似的赋值效果
func variable6(){
	var i, j, k, l = 3, 9, true, "frank"
	var s = "Jefflikie"
	fmt.Println(i, j, s, k, l)// 3 9 Jefflikie true frank
}

// 省去var关键字,:=的意思就是，定义并声明变量
func variable7(){
	i, j, k, l := 3, 9, true, "frank"
	// 既然:=是定义并赋值，那么我们后面再进行定义赋值就是重复定义了，所以会报错
	//i := 5
	i = 5
	s := "Jefflikie"
	fmt.Println(i, j, s, k, l)// 5 9 Jefflikie true frank
}

//定义一个包内部的变量
var pa = 123


func main() {
	variable1()
	variable2()
	variable3()
	variable4()
	variable5()
	variable6()
	variable7()
	fmt.Println(pa)//123
}
