package main

import "fmt"

func test1()  {
	fmt.Printf("hello jefflike\n")
}

func test2() int {
	return 1
}

func test3(a,b,c int) (int,int,int) {
	return a,b,c
}

func test4() (a int,b int,c int) {
	a = 1
	b = 2
	c = 3
	return
}

func main()  {
	test1()
	//a:=test2()
	//fmt.Println(a)

	//a,b,c :=test3(1,2,3)
	//fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
	a,b,c := test4()
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}
