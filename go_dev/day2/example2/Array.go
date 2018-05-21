/*
go语言的数组是值传递类型
go元使用数组时的函数里的形参必须是固定长度的，这样使用起来是很不方便的
我们使用go元的数组通常是使用数组的切片

*/
package main

import "fmt"

func array1()  {
	// 定义一个数组
	var arr1 [5]int //[0 0 0 0 0]
	//var arr1 = [5]int{1,2,3,4,5}

	// :=要给数组赋初值
	arr2 := [3]int{1,2,3}//[1 2 3]
	fmt.Println(arr1, arr2)

	//奇葩的赋值方式
	arr3 := [...]int{2,4,6,7,10}//[2 4 6 7 10]
	fmt.Println(arr3)

	//二维数组的定义
	var grid [4][5]int//[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	fmt.Println(grid)

	//数组的遍历
	for i:=0;i<len(arr2);i++{
		fmt.Println(arr2[i])//1   2    3
	}

	//range获取数组的下标，遍历数组
	for i:=range arr2{
		fmt.Println(arr2[i])//1   2    3
	}

	//range还可以获得数组的值
	for i, v:=range arr2{
		fmt.Println(i, v)//1   2    3
	}

	//只要v的情况
	sum := 0
	for _, v := range arr2{
		sum += v
	}
}

//go语言的数组是值类型的，也就是说，在函数内是不会改变函数外调用的数组的值的
func array2( arr [3]int){
	arr[0] = 100
	for _,v :=range arr{
		fmt.Println(v)//100,3,4,此处相当于将复制得来的数组改变了，但是外层的数组并没有改变
	}
}

// 取指针
func array3( arr *[3]int){
	arr[0] = 100
	for _,v :=range arr{
		fmt.Println(v)//100,3,4,此处相当于将复制得来的数组改变了，但是外层的数组并没有改变
	}
}

func main() {
	//var arr1 [5]int
	arr2 := [3]int{2,3,4}
	array1()
	//array2(arr1)//这里是会报错的，array2函数需要传入一个长度为3的数组，go语言里认为长度为3和长度为5的数组不是一个类型的，所以会飘红
	array2(arr2)
	fmt.Println(arr2)//[2 3 4]

	//array3接收的是一个数组的指针，所以我们传一个arr的地址值进去
	array3(&arr2)
	fmt.Println(arr2)//[100 3 4]，此时是相当于引用传递的，所以会改变数组的值
}
