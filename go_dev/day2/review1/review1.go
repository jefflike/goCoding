package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
)

//整行文件的读写
func fileIo(){
	const filename = "a.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else{
		//将文件内容格式化输出
		fmt.Println(string(contents))
		// 打印文本文件要使用格式化输出
		fmt.Printf("%s \n",contents)
	}
}

//if的另一种操作方式
func fileIo1(){
    const filename = "a.txt"
    contents, err := os.Open(filename)//&{0xc042048180},这里的contents类似一个文件句柄，是一个指针的地址

    if err != nil{
    	panic(err)
	}else{
		scanner := bufio.NewScanner(contents)
		for scanner.Scan(){
			fmt.Println(scanner.Text())
	}

	}

    fmt.Println(contents,err)//&{0xc042048180} <nil>
}

//常量与枚举
func myConst(){
	const(
		a = iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d )
}


func main() {
	// os.Open
	fileIo1()

	//整篇打印出文章
	fileIo()

	//打印枚举类型
	myConst()

    // 打印hello world
	fmt.Println("hello")
}
