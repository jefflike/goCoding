/*
go的函数是编程，函数是一等公民
*/
package main

import "fmt"

func operation(op func(int, int) int,a, b int) int {
    return op(a, b)
}

func add(m, n int) int {
	return m + n
}

func main() {
	v := operation(add, 5, 6)
	fmt.Println(v)
}
