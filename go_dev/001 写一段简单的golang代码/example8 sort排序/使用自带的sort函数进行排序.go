/*
排序整型使用sort.Ints
*/
package main

import (
	"sort"
	"fmt"
)

func main()  {
	a := []int{3, 5, 1, 43, 9, 34}
	sort.Ints(a)
	for i, v := range a{
		fmt.Printf("第%d个是%d\n", i, v)
	}
}