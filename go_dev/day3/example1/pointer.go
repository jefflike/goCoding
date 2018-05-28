/*
&a是取变量的地址
不要指向没有合法指向的内存
*/
package main

import "fmt"

func pointer()  {
	a := 10
	fmt.Println(a, &a)//10 0xc0420361d0

	var p *int
	fmt.Println(p)//<nil>

	p = &a//p的值取的就是a的地址
	*p = 123//指针指到p所指的内存，并赋值
	fmt.Println(a)//123

	//*p直接指向一个值，这个值是不确定的，没有具体的内存空间，所以指向不能成功
	//*p = a//panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(p)
}

func pointer2(){
	a := 10
	var p *int
	// 指针其他的指向内存的方式,动态分配空间（但是只申请不需要手动释放）
	p = new(int)//相当于为指针p，new一个int大小的内存块用于存数据
	p = &a
	fmt.Println(*p)//10

	q := new(string)
	*q = "jefflike"
	fmt.Println(*q)
}

func main() {
	//pointer()
	pointer2()
}
