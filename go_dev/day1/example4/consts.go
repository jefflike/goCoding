/*
const定义常量可以不指定常量的数据类型，
那么在表达式中使用到常量的过程类似于赋值的过程

go使用const常量的方式表示枚举
*/

package main

import (
	"math"
	"fmt"
)

// 定义多个常量,const必须要赋值
const(
	m, n = "1", 5
	o = true
	p = 3
)

func trianagle(){
	const a, b = 3, 4
	//const a, b int = 3, 4
	var c int
	// a,b是常量，ab的类型是由表达式而改变的，在此处a，b先带入表达式，因为需要为float
	// 所以结果就被float类型接收了
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(c)
}

// 枚举
func enums(){
	//iota是一种枚举递增的意思
	const(
		first = iota
		second
		third
		fourth
	)
	fmt.Println(first, second, third, fourth) // 0 1 2 3
}

//定义b,kb,mb,gb,tb,pb
func save()  {
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824 1099511627776 1125899906842624
}

func main() {
	trianagle()
	fmt.Println(m, n, o, p) //1 5 true 3
	enums()
	save()
}
