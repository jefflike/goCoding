package main

import (
	"go_dev/pipeline"
	"fmt"
)

func main() {
	p := pipeline.Merge(
		pipeline.IntMemorySort(pipeline.ArraySource(3,6,8,2,5)),
		pipeline.IntMemorySort(pipeline.ArraySource(34,2,4,5)),
	)
	// 因为我的pipeline明确标明了close所以我可以使用简单的遍历方式
	//for {
	//	if v,ok := <- p; ok{
	//		fmt.Println(v)
	//	}else {
	//		break
	//	}
	//}

	for v := range p{
		fmt.Println(v)
	}
}
