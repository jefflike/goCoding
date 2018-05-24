package main

import "fmt"

func main() {
	a := 10
	b := "jefflike"

	//无参无返回值闭包1
	f1 := func (){
		fmt.Printf("a=%d,b=%s\n",a,b)
	}
	f1()

	func(){
		fmt.Printf("a=%d,b=%s\n",a,b)
	}()

	//有参无返回值1
	f2 := func (i,j int){
		fmt.Printf("a=%d,b=%d\n",i,j)
	}
	f2(1,2)

	//有参无返回值2
	func (i,j int){
		fmt.Printf("a=%d,b=%d\n",i,j)
	}(1,2)

	//有参有返回值1
	f3 := func (i,j int)(result int){
		result = i+j
		return
	}
	x := f3(1,2)
	fmt.Println(x)

	//有参有返回值2
	y := func (i,j int)(result int){
		result = i+j
		return
	}(1,2)
	fmt.Println(y)

	//闭包函数内部是同一个作用域
	//此时	a := 10 b := "jefflike"
	func(){
		a = 20
		b= "frank"
		fmt.Printf("内部a=%d,b=%s\n",a,b)//内部a=20,b=frank
	}()
	fmt.Printf("外部a=%d,b=%s",a,b)// 外部a=20,b=frank
}
