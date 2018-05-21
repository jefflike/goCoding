package main

import "fmt"

func main() {
	// go starts a goroutine,他的功能类似于python的thread，新开一个线程的操作类似
	// goroutine 的功能远强于thread
	// 此时运行程序的结果是什么都没有显示，因为执行main函数的同时执行了goroutine，在打印操作之前
	// goroutine的操作已经结束了，那么main函数执行结束，程序退出，所以不会执行打印hello的操作。
	go printHello()
}

func printHello(){
	fmt.Println("hello")
}
