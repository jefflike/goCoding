package main

import "fmt"

func more(s string) int {
	start := 0 //当前找到最长不含重复字符串的起始位置
	maxLenght := 0 //最长字符串的长度
	lastOccurred := make(map[byte] int)//记录每个字母最后出现的位置（存到一个map里）
	for i, v:=range []byte(s) {
		lastI, ok := lastOccurred[v]
		if ok && lastI >= start {start = lastI +1}
		if i-start+1 >maxLenght{maxLenght = i -start +1}
		lastOccurred[v] = i
	}
	return maxLenght
}

func main(){
	fmt.Println(more("dajuidhgaui"))
}
