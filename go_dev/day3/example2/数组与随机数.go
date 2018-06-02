/*
所谓的数组比较，其实只支持==或!=,两个数组要进行比较需要确定数据类型长度都相等才能比
切片也是一种数据结构，引用之前数组的长度等
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func compare1()  {
	a := [5]int{1,2,3,4,5}
	b := [...]int{1,2,3,4,5}
	c := [5]int{1,2,3}
	d := [5]int{3,2,1}
	fmt.Println("a==b:", a==b)//a==b: true
	fmt.Println("a==c:", a==c)//a==c: false
	fmt.Println("d==c:", d==c)//d==c: false

	e := d
	fmt.Println(e)//[3 2 1 0 0]
}

func random1()  {
	// 多循环几次会发现结果都是一样的，
	// 想获得随机值需要向seed传入改变的值，比如当前时间
	rand.Seed(123)
	for i :=0; i<10; i++{
		fmt.Println(rand.Int())
	}
}

func random2()  {
	// 现在随机数会发生改变了，但是数值非常的大
	rand.Seed(time.Now().UnixNano())
	for i :=0; i<10; i++{
		fmt.Println(rand.Int())
	}
}

func lenCap()  {
	a := []int{1,2,3,0,0}
	s := a[0:3:5]
	fmt.Println("s=",s)//s= [1 2 3]
	fmt.Println("len(s):",len(s))//len(s): 3
	fmt.Println("cap(s)",cap(s))//cap(s) 5
}

func random3()  {
	// 限制随机数的大小
	for i :=0; i<10; i++{
		fmt.Println(rand.Intn(1000))
	}
}

//数组的长度不会变化，但是slice的长度可以变化，变化的长度是2的n次方（最小的且能装下所有的slice数据的最小长度）
func lenCap2()  {
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d",len(a),cap(a))//5,5

	b := []int{}
	fmt.Printf("len = %d, cap = %d",len(b),cap(b))//0,0

	b = append(b, 1,2,3)
	fmt.Printf("len = %d, cap = %d",len(b),cap(b))//3,4
}

func main() {
	//compare1()
	//random1()
	//random2()
	//random3()
	//lenCap()
	lenCap2()
}
