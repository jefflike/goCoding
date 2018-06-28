/*
go语言的切片与python的十分相似的
go文档里说，slice是对数组的一个视图（view），只看我要看的这一小段
go的切片也是半开半闭的
slice本身是没有数据的，是对底层array的一个view,所以改动slice就是等于改动了底层的array
reslice概念
extending slice概念：向后扩展但不超过cap(),末尾不能超过len()
python的列表切片没有cap()的概念，切不到的都没有，因为他是一个引用的类型
*/
package main

import "fmt"

// 中括号不加数字就是一个切片
func slice1(arr []int){
	//改动切片的数值
	arr[0] = 100
}

func printSlice(a []int)  {
	fmt.Printf("a=%v, len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
}

func main() {
	//arr := [5]int{2,4,6,8,10}
	arr := [...]int{2,4,6,8,10}
	//两端都有
	s := arr[1:4]
	fmt.Println(s)//[4 6 8]
	//有前无后
	fmt.Println(arr[2:])//[6 8 10]
	//有后无前
	fmt.Println(arr[:4])//[2 4 6 8]
	//全切
	fmt.Println(arr[:])//[2 4 6 8 10]

	slice1(arr[:4])
	fmt.Println(arr)//[100 4 6 8 10]

	//reslice操作,每次的下标都是针对自己的slice
	a := arr[:]
	a = a[:3]
	a = a[1:]
	fmt.Println(a)//[4,6]

	// extending slice
	arr1 := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr1[2:6]
	fmt.Printf("s1=%v cap=%d,len=%d\n", s1, cap(s1), len(s1))//s1=[2 3 4 5] cap=6,len=4
	//很明显此时s1的长度是4，即下标是0-3，那么此时的s2会不会报错，答案是不会报错
	s2 := s1[3:5]
	fmt.Println(s2)//[5 6]
	fmt.Printf("s2=%v cap=%d,len=%d\n", s2, cap(s2), len(s2))//s2=[5 6] cap=3,len=2
	s3 := s1[3:]
	fmt.Println(s3)//[5],默认是看不到cap()的值的

	//此处修改了原数组arr1
	s4:=append(s2,10)
	//此处view的数组i已经不是原始的arr1了，系统为我们分配了一个更加长的数组进行view
	s5 := append(s4,11)
	s6 :=append(s5,12)
	fmt.Println("s4",s4)// s4 [5 6 10]
	fmt.Println("s5",s5)// s5 [5 6 10 11]
	fmt.Println("s6",s6)// s6 [5 6 10 11 12]
	fmt.Println("arr1",arr1)//arr1 [0 1 2 3 4 5 6 10]
	fmt.Println(cap(s4))// 3
	fmt.Println(cap(s5))// 6
	fmt.Println(cap(s6))// 6

	var b []int
	for i:=1;i<100;i++{
		b = append(b, i)
		printSlice(b)
	}
	fmt.Println(b)

	ab := []int{1,2,3,4,5,6,7,8,9}
	// 删除第三个数
	//ab = append(ab[:2], ab[3:]...)
	//fmt.Println(ab)

	//head := ab[0]
	//ab = ab[1:]
	//fmt.Println(head)
	//fmt.Println(ab)

	tail := ab[len(ab)-1]
	ab = ab[:len(ab)-1]
	fmt.Println(tail)
	fmt.Println(ab)

}
