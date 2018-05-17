package main

import(
	"go_dev/day1/package_example/calc"
	"fmt"
)

func main(){
	sum := calc.Add(5, 8)
	fmt.Println(sum)
}