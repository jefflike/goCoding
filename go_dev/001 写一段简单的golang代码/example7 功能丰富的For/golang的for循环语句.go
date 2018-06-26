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

func tes(){
	s := 0
	for i :=1; i<100; i++{
		if i%2==0{continue}
		s += i;

	}
	fmt.Println(s)
}
func tes1()  {
	var a [5]int = [5]int{5,7,3,9,2}

	for i:=0; i< len(a)-1; i++{
		for j := 0;j < len(a) -i -1;j++  {
			if a[j]>a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	for s :=0; s< len(a);s++{
		fmt.Println(a[s])
	}
}

func DiGui( a int) int {
	if a == 1{
		return 1
	}else {
		return a + DiGui(a-1)
	}

}

func diGui(a int) int  {
	if(a == 1){
		return 1
	}
	return a * diGui(a-1)
}

func main() {
	sum := 0
	for i :=1; i<=20; i++ {
		sum += diGui(i)
	}
	fmt.Println(sum)

	//fmt.Println(DiGui(100))

	tes1()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
	printFile("a.txt")
}
