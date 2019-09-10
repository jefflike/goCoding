/*
go的内建变量的类型
    bool，string
    每一种int的类型都可以分为有无符号的整形，加u就是无符号整型
    int8-64，规定了长度，int不规定长度，那么在32位系统的长度就是32位，64位系统为64位
    uintptr（指针）
    (u)int, int8, ingt16, int32, int64, uintptr

    rune是go语言的字符型，但是不同于其它语言的char，rune长度是32位的（4个字节），其实在现在的全球化(汉字单个字utf8是三个字节（24位），Unicode都是两个字节(16位))
	一个字符的char类型已经是一个坑了，所以rune更加好用。
	byte就是一个字节长度的8位
    byte，rune

	go语言里没有double的类型，float32, float64作为浮点数类型
    complex64,complex128 很少见的把复数类型作为了内建的变量类型
	float32, float64, complex64,complex128,

    强制类型转换，
	go语言里只有强制类型转换，没有隐式类型转换，如变量提升这种
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//定义一个复数，求模长
func comp() {
	//这样写就是一个复数
	c := 3 + 4i
	length := cmplx.Abs(c)
	fmt.Println(length) // 5
}

//使用复数验证一下欧拉函数
func euler() {
	//欧拉公式
	value := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(value) //(0+1.2246467991473532e-16i)
}

/*
??欧拉公式的结果不是0？，我们的complex的类型是complex64与128，他的实部与虚部分别是32位和
64位的float,float浮点型都是不够精准的
*/

// 强制类型转换的小例子
func trianagle() {
	a, b := 3, 4
	var c int
	// math.Sqrt要求参数是一个浮点类型，但是我们的a与b都是整型，不能进行自动的类型转换
	//需要我们手动进行类型转换
	//同理我们math.Sqrt的结果也是浮点型，我们定义了int型的c接收就必须强制类型转换
	//c = int(math.Sqrt(a * a + b * b))
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 键盘输入的两种方式
func input1() {
	var a int // 声明变量a
	fmt.Println("请输入:")
	fmt.Scanf("%d", &a)
	fmt.Println("a is ", a)
}

func input2() {
	var b int // 声明变量a
	fmt.Println("请输入:")
	fmt.Scan(&b)
	fmt.Println("b is ", b)
}

func main() {
	comp()
	euler()
	a := "我爱你北京"
	fmt.Println(len(a)) //15
	//input1()
	input2()
}
