/*
go 的switch..case语句是不需要break的，默认每个case后面就有break
fallthrough是用于case后面不加break的时候用的。

go里的switch的case里可以放逻辑表达式
panic的作用类似与python的raise
*/
package main

import (
	"fmt"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		// 报错，让程序停下来
		panic("非法运算符：" + op)
	}
    return result
}

//switch没有表达式，判断部分就发给到case里进行
func testswitch(score int) string {
	g := ""

	switch {
	case score>100 || score<0:
		panic("非法成绩")
	case score > 90 :
		g = "profact"
	case score > 80:
		g = "good"
	case score > 60:
		g = "just so so"
	default:
		g = "fuck"
	}
	return g
}

func main() {
	v := eval(12, 6, "/")
	fmt.Println(v)

	s := testswitch(45)
	fmt.Println(s)// fuck

	s = testswitch(123)
	fmt.Println(s)// 抛异常，但是依然会打印出上面正确的成绩
}
