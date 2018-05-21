/*
go语言的数组是值类型
*/
package main

import "fmt"

func array1()  {
	// 定义一个数组
	var arr1 [5]int//[0 0 0 0 0]

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

//go语言的数组是值类型的
func array2( arr [3]int){
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
	fmt.Println(arr2)
}
