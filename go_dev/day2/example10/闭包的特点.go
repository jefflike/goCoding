package main

import "fmt"

//一般函数
func test() int {
	//函数被调用时a被分配一块内存空间
	var a int//初始化的a为0值
	a++//a=1
	return a*a//函数调用完毕，a被释放
}

//闭包函数
func biBao() func() int{
	//闭包函数不关心这些捕获的变量是否超过了作用域，只要闭包还在使用，这些变量就还会存在
	var a int
	return func() int {
		a++
		return a * a
	}
}

func main() {
	fmt.Println(test())//1
	fmt.Println(test())//1
	fmt.Println(test())//1
	f := biBao()
    fmt.Println(f())//1
    fmt.Println(f())//4
    fmt.Println(f())//9
}
