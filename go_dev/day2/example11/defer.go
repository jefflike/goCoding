/*
gopath环境变量，配置为src同级目录的绝对路径
想要自动生成bin和pkg需要使用go install命令。
除了要配置gopath还需要配置gobin。
这三个目录是用来做什么的
src：放源代码的目录
bin：放可执行程序的目录
pkg：放平台相关的库
*/
package main

import "fmt"

//执行这个文件最开始就会运行这个函数（不需要调用自动运行）
//导入其他包时，如果其他包有init函数会在导入的时候执行
//这一点与python的init文件在导包的执行情况比较相似
func init(){
	fmt.Println("初始化执行init")
}

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
