/*
　　go语言设计的初衷。

　　1.针对其他语言的痛点进行设计。

　　2.并且加入了原生并发的设计，即加入了并发编程。

　　3.为大数据，微服务并发而生的通用编程语言。(即与c，c++，python的语言一样可以完美运行)。

　　go语言的特别之处

　　1.与其他的面向对象的编程语言不同的地方是，go语言没有"对象"，没有继承多态，没有泛型，没有try/catch.（没有class）

　　2.go语言有什么：go语言有接口，函数式编程，csp并发模型（goroutine+channel）。
*/
package main

import(
	"fmt"
)

func main(){
	fmt.Println("hello world!")
}