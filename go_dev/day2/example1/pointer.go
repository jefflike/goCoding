/*go中的指针是一个简单的概念
go中的指针不可以运算，不像c语言的指针拿到首地址可以i一直++
go语言是值传递语言，即形参是值的copy赋值而不是引用类型传递
准确的说go语言的参数传递是值传递加指针传递结合使用的。
对于go语言的自定义对象类型，我们可以手动的将类的值类型写成值传递方式或者指针传递方式。
*/
package main

import "fmt"


//go语言的值传递与python的值传递方式有些相似
//但是在此处改变的其实是值传递的值，但是并没有改变实际的ab的值，即此处是值传递，不是引用传递
func exchange(a,b int){
	a,b = b,a
	fmt.Println("a的值：",a,",b的值：",b)
}

// 指针传递，引用传递,但是我们一般不这么用
func exchange1(a,b *int){
	*a,*b = *b,*a
	fmt.Println("a的值：",a,",b的值：",b)
}

// 一般转换两个数的值会使用这种方式
func exchange2(a,b int)(int, int){
	return b, a
}

func main() {
	var a = 2
	var pa *int = &a
	*pa = 3
	//*(&a) = 10
	fmt.Println(a)//3
	b :=13
	exchange(a,b)
	fmt.Println("a的值：",a,",b的值：",b)

	exchange1(&a,&b)
	fmt.Println("a的值：",a,",b的值：",b)

	exchange2(a, b)
	fmt.Println("a的值：",a,",b的值：",b)
}
