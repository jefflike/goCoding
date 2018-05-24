package main

import "fmt"

func defer1(){
	a := 10
	b := 20
	defer func(){
		fmt.Printf("内部a=%d,b=%d\n",a,b)
	}()

	a = 20
	b = 10
	fmt.Printf("外部a=%d,b=%d\n",a,b)
	//外部a=20,b=10
	//内部a=20,b=10
}

func defer2(){
	a := 10
	b := 20
	defer func(a,b int){
		fmt.Printf("内部a=%d,b=%d\n",a,b)
	}(a,b)//这一步的时候，实际a，b的值已经传递进去了，所以打印还是10，20

	a = 20
	b = 10
	fmt.Printf("外部a=%d,b=%d\n",a,b)
	//外部a=20,b=10
	//内部a=10,b=20
}

func main() {
    // defer在main函数结束之前的一瞬间结束一些文件读写，网络连接的操作，相当于python上下文管理器的__exit__()
	defer fmt.Println("bbbbbbbbbbbbb")
	fmt.Println("aaaaaaaa")

	defer1()
	defer2()
}
