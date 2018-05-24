package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入一个整型变量")
	// scan输入
	fmt.Scanf("%d", &a)
	fmt.Println(a)//输入的如果不是一个整形，得到的结果就是0

	// scan输入2
	fmt.Scan(&a)
	fmt.Println(a)
}
