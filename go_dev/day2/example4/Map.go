/*
go语言不用担心变量不初始化的问题，即使没有初始化也不会报错，而是使用zerovalue
map的遍历与数组的遍历方式是一样的
*/
package main

import "fmt"

// map的定义1
func map1(){
	m := map[string]int{
		"k1": 1,
		"k2": 2,
	}
	fmt.Println(m)// map[k1:1 k2:2]
}
//map的定义2
func map2()  {
	m := make(map[string]string) // 这里定义的是一个empty map
	m["k1"] = "j"
	m["k2"] = "e"
	m["k3"] = "f"
	m["k4"] = "f"
	fmt.Println(m)//map[k1:j k2:e k3:f k4:f]
	fmt.Println(m["xixi"])//没有这个值，会打印一个空串就是我们所说的zerovalue

	//python里in的判断操作
	if v,tr := m["x"];tr{
		fmt.Println(v)//tr是true就打印v
	}else {
		fmt.Println("key don't exist")
	}

	//pop值
	delete(m,"k4")
	fmt.Println(m)
}

// map定义3
func map3(){
	var m map[string]string//这里定义的是一个nil
	//m["k1"] = "j"//此时这样赋值是不成立的，这个m此时是nil
	fmt.Println(m)
}

func main() {
	map1()
	map2()
	map3()
}
