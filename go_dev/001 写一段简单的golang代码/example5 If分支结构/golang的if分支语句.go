/*
go的条件判断语句中，if后面是不需要加括号的
go的函数的返回值可以返回两个值
go打印可以使用ioutil.ReadFile(),一行一行的打印可以使用Os.oppen(文件)，a := bufio.NewScaner(contents)
遍历a.Scan(),打印a.Text()就是遍历打印出文件的数据
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func testIf(a int) int {
	if a > 100{
		return 100
	}else if a < 0{
		return 0
	}else {
		return a
	}
}

// 文件读写
func fileIo(){
	const filename = "a1.txt"
	//go的读写会返回两个值，读到的内容和err，用两个值来接收
	contents, err := ioutil.ReadFile(filename)
	// 判断是不是是为空
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}

}

// if的简写
func fileIo2(){
	const filename = "a.txt"
	//这里的contents的名称空间只在if的里面，出了if就访问不到了
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}

}

func main() {
	v := testIf(14)
	fmt.Println(v) //14
	fileIo() //open a.txt: The system cannot find the file specified.
	// 如果有文件，那么久打开文件的内容
	fileIo2() //hello world
}
