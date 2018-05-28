/*
for循环的循环条件同样不需要加括号
for的三个表达式可以省略
go里面是没有while循环的，因为for已经包含了while的功能，没必要再多一种语法
go里转换成string类型的方法是使用strconv.Iota(int转换成字符串类型)
*/
package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

//没有起始值的for循环
func convertToBin(n int) string {
	result := ""
	for  ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//只有终止值的for循环
func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil{
		panic("文件异常")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//while(true),就是死循环，会时常使用
func forever(){
	for{
		fmt.Println("...")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
	printFile("a.txt")
}
