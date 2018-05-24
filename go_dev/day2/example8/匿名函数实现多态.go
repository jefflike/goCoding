package main

import "fmt"

type myFunc func(int, int) int

func Calc(a,b int, cal myFunc) (result int) {
	result = cal(a,b)
	return
}

func Add(a,b int) int {
	return a+b
}

func Min(a,b int) int {
	return a-b
}

func main() {
	v1 := Calc(34, 12, Add)
	v2 := Calc(34, 12, Min)
	fmt.Printf("v1=%d,v2=%d",v1,v2)//v1=46,v2=22
}
